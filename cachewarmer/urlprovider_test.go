package cachewarmer

import "testing"

// mock url provider is defined in ./cachewarmer_test.go

func TestUrlProviderExists(t *testing.T) {
	if !urlProviderExists("sitemap") {
		t.Fail()
	}

	if urlProviderExists("testa") {
		t.Fail()
	}
}

func TestUrlProviderKeys(t *testing.T) {
	keys := urlProviderKeys()

	if len(keys) != 1 {
		t.Fail()
	}

	if keys[0] != "sitemap" {
		t.Fail()
	}
}

func TestGetUrlProviderByName(t *testing.T) {
	if p, _ := getUrlProviderByName("sitemap"); p == nil {
		t.Fail()
	}

	if p, _ := getUrlProviderByName("testa"); p != nil {
		t.Fail()
	}
}
