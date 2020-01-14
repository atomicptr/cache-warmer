package cachewarmer

import (
	"strings"
	"testing"
)

func TestApplyPrefixUrl(t *testing.T) {
	urls := []string{
		"https://example.com/page-1",
		"https://example.com/page-2",
		"https://example.com/page-3",
		"https://example.com/page-4",
		"https://example.com/page-5",
		"https://example.co.uk/page-1",
		"https://example.super.long.suburl.domain.com/page-1",
	}

	prefixUrl := "https://atomicptr.de"
	urls = applyPrefixUrl(prefixUrl, urls)

	for _, url := range urls {
		if !strings.HasPrefix(url, prefixUrl) {
			t.Errorf("expected url %s to start with %s", url, prefixUrl)
		}
	}
}
