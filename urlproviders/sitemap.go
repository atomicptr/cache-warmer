package urlproviders

import (
	"errors"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/atomicptr/cache-warmer/cachewarmer"
	"github.com/beevik/etree"
)

type SitemapUrlProvider struct {
}

func init() {
	cachewarmer.AddUrlProvider("sitemap", &SitemapUrlProvider{})
}

func (s *SitemapUrlProvider) FetchUrls(path string, client *http.Client) ([]string, error) {
	var urls []string
	var xmlDataBlob io.Reader
	var err error

	if strings.HasPrefix(path, "http") {
		xmlDataBlob, err = s.fetchXmlFromWeb(path, client)
	} else {
		xmlDataBlob, err = s.fetchXmlFromFile(path)
	}

	if err != nil {
		return nil, err
	}

	doc := etree.NewDocument()
	if _, err := doc.ReadFrom(xmlDataBlob); err != nil {
		return nil, err
	}

	// check if the sitemap is a sitemap index
	sitemapIndex := doc.FindElement("sitemapindex")

	if sitemapIndex != nil {
		for _, sitemap := range sitemapIndex.ChildElements() {
			loc := sitemap.FindElement("loc")
			if loc != nil {
				sitemapUrls, err := s.FetchUrls(loc.Text(), client)
				if err != nil {
					return nil, err
				}
				urls = append(urls, sitemapUrls...)
			}
		}
	}

	// regular sitemap
	urlset := doc.FindElement("urlset")

	if urlset != nil {
		for _, url := range urlset.ChildElements() {
			loc := url.FindElement("loc")
			if loc != nil {
				url := loc.Text()
				urls = append(urls, url)
			}
		}
	}

	return urls, nil
}

func (s *SitemapUrlProvider) fetchXmlFromWeb(path string, client *http.Client) (io.Reader, error) {
	resp, err := client.Get(path)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	return resp.Body, nil
}

func (s *SitemapUrlProvider) fetchXmlFromFile(path string) (io.Reader, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return file, nil
}
