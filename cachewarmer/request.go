package cachewarmer

import (
	"net/http"
)

func (cw *CacheWarmer) makeRequestsFromUrls(urls []string) ([]*http.Request, error) {
	var requests []*http.Request

	for _, url := range urls {
		req, err := cw.makeRequestFromUrl(url)
		if err != nil {
			return nil, err
		}

		requests = append(requests, req)
	}

	return requests, nil
}

func (cw *CacheWarmer) makeRequestFromUrl(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", UserAgent)

	for headerKey, headerValue := range cw.config.HeaderMap() {
		req.Header.Set(headerKey, headerValue)
	}

	for cookieKey, cookieValue := range cw.config.CookieMap() {
		req.AddCookie(&http.Cookie{
			Name:  cookieKey,
			Value: cookieValue,
		})
	}

	return req, nil
}
