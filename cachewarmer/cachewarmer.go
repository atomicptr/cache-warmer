package cachewarmer

import (
	"log"
	"net/http"
	"net/url"
)

const UserAgent = "atomicptr/cache-warmer"

type CacheWarmer struct {
	config Config
	client http.Client
	logger *log.Logger
}

func New(config Config, logger *log.Logger) (*CacheWarmer, error) {
	err := config.Validate()
	if err != nil {
		return nil, err
	}

	return &CacheWarmer{
		config: config,
		client: http.Client{
			Timeout: config.HttpClientTimeout,
		},
		logger: logger,
	}, nil
}

func (cw *CacheWarmer) Run() error {
	provider, err := getUrlProviderByName(cw.config.Provider)
	if err != nil {
		return err
	}

	urls, err := provider.FetchUrls(cw.config.Path, &cw.client)
	if err != nil {
		return err
	}

	cw.logger.Printf("%d urls found...", len(urls))

	urls = applyPrefixUrl(cw.config.PrefixUrl, urls)

	requests, err := cw.makeRequestsFromUrls(urls)
	if err != nil {
		return err
	}

	crawler := Crawler{
		Client:          cw.client,
		NumberOfWorkers: cw.config.NumberOfWorkers,
		Logger:          cw.logger,
	}

	crawler.Crawl(requests)

	return nil
}

func applyPrefixUrl(prefixUrl string, urls []string) []string {
	if prefixUrl == "" {
		return urls
	}

	prefixUrlParsed, err := url.Parse(prefixUrl)
	if err != nil {
		return urls
	}

	var newUrls []string

	for _, rawUrl := range urls {
		u, err := url.Parse(rawUrl)
		if err == nil {
			u.Scheme = prefixUrlParsed.Scheme
			u.Host = prefixUrlParsed.Host

			newUrls = append(newUrls, u.String())
		}
	}

	return newUrls
}
