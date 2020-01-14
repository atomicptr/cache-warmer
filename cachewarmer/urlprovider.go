package cachewarmer

import (
	"fmt"
	"net/http"
)

var urlProviderList map[string]UrlProvider

type UrlProvider interface {
	FetchUrls(path string, client *http.Client) ([]string, error)
}

func AddUrlProvider(name string, provider UrlProvider) {
	if urlProviderList == nil {
		urlProviderList = make(map[string]UrlProvider)
	}

	urlProviderList[name] = provider
}

func urlProviderExists(name string) bool {
	_, ok := urlProviderList[name]
	return ok
}

func urlProviderKeys() []string {
	var keys []string
	for key := range urlProviderList {
		keys = append(keys, key)
	}
	return keys
}

func getUrlProviderByName(name string) (UrlProvider, error) {
	if provider, ok := urlProviderList[name]; ok {
		return provider, nil
	}

	return nil, fmt.Errorf("unknown url provider %s", name)
}
