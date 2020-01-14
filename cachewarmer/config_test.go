package cachewarmer

import (
	"net/http"
	"os"
	"strings"
	"testing"
)

type mockSitemapProvider struct {
}

func (m mockSitemapProvider) FetchUrls(_ string, _ *http.Client) ([]string, error) {
	return nil, nil
}

func TestMain(m *testing.M) {
	AddUrlProvider("sitemap", mockSitemapProvider{})
	os.Exit(m.Run())
}

func TestValidateWithInvalidProvider(t *testing.T) {
	c := Config{
		Provider: "Invalid Provider",
	}

	err := c.Validate()

	if err == nil {
		t.Error("Validation should have failed, because the provider is invalid")
	}
}

func TestValidateWithInvalidPrefixUrl(t *testing.T) {
	c := Config{
		Provider:  "sitemap",
		Path:      "https://example.com/sitemap.xml",
		PrefixUrl: "this is not a url!",
	}

	err := c.Validate()

	if err == nil {
		t.Error("Validation should have failed, because the prefix url is invalid")
	}
}

func TestValidate(t *testing.T) {
	c := Config{
		Provider:  "sitemap",
		Path:      "https://example.com/sitemap.xml",
		PrefixUrl: "https://atomicptr.de",
	}

	err := c.Validate()

	if err != nil {
		t.Fail()
	}
}

func TestValidateKeyValueSetWithValidInputs(t *testing.T) {
	kvSet := []string{
		"a=b",
		"b=5",
		"test_name=value-that-is-kinda-long",
	}

	err := validateKeyValueSet("test", kvSet)

	if err != nil {
		t.Error(err)
	}
}

func TestValidateKeyValueSetWithInvalidInputs(t *testing.T) {
	kvSet := []string{
		"a=b",
		"b=5",
		"lorem ipsum dolor sit amet",
	}

	err := validateKeyValueSet("test", kvSet)

	if err == nil {
		t.Fail()
	}
}

func TestCreateMapFromKeyValueStrings(t *testing.T) {
	kvSet := []string{
		"a=b",
		"b=5",
		"test_name=value-that-is-kinda-long",
	}

	kvMap := createMapFromKeyValueStrings(kvSet)

	for _, kvPair := range kvSet {
		parts := strings.Split(kvPair, "=")

		if kvMap[parts[0]] != parts[1] {
			t.Errorf("Expected '%s' to be '%s'", parts[0], parts[1])
		}
	}
}
