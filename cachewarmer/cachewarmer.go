package cachewarmer

import (
	"net/http"
)

type CacheWarmer struct {
	config Config
	client http.Client
}

func New(config Config) (*CacheWarmer, error) {
	err := config.Validate()
	if err != nil {
		return nil, err
	}

	return &CacheWarmer{
		config: config,
		client: http.Client{
			Timeout: config.HttpClientTimeout,
		},
	}, nil
}

func (cw *CacheWarmer) Run() error {
	// TODO: "fetch" paths
	// TODO: prepare requests
	// TODO: execute requests
	return nil
}

func (cw *CacheWarmer) Stop() {
	// TODO: stop everything
}
