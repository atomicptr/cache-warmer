package cachewarmer

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

type Config struct {
	Provider          string
	Path              string
	PrefixUrl         string
	Cookies           []string
	Headers           []string
	HttpClientTimeout time.Duration
	NumberOfWorkers   int

	cookieMap map[string]string
	headerMap map[string]string
}

// Checks if the configuration is valid, returns error if invalid.
func (c *Config) Validate() error {
	if !urlProviderExists(c.Provider) {
		return fmt.Errorf(
			"provider is invalid, should be one of the following: %s",
			strings.Join(urlProviderKeys(), ", "),
		)
	}

	// if it doesn't start with http it's probably a file path, check if that file exists...
	if !strings.HasPrefix(c.Path, "http") {
		_, err := os.Stat(c.Path)
		if os.IsNotExist(err) {
			return fmt.Errorf("file %s does not exist", c.Path)
		}
	}

	if c.PrefixUrl != "" && !strings.HasPrefix(c.PrefixUrl, "http") {
		return fmt.Errorf("prefix url is not a proper url: %s", c.PrefixUrl)
	}

	err := validateKeyValueSet("cookie", c.Cookies)
	if err != nil {
		return err
	}

	err = validateKeyValueSet("header", c.Headers)
	if err != nil {
		return err
	}

	return nil
}

func validateKeyValueSet(name string, keyValueSet []string) error {
	regex, err := regexp.Compile(`.+=.+`)
	if err != nil {
		return err
	}

	for _, keyValuePair := range keyValueSet {
		if regex.MatchString(keyValuePair) {
			return fmt.Errorf(
				"%s does not match pattern %s_name=%s_value for: %s",
				name,
				name,
				name,
				keyValuePair,
			)
		}
	}
	return nil
}

func (c *Config) HeaderMap() map[string]string {
	if c.headerMap == nil {
		c.headerMap = createMapFromKeyValueStrings(c.Headers)
	}
	return c.headerMap
}

func (c *Config) CookieMap() map[string]string {
	if c.cookieMap == nil {
		c.cookieMap = createMapFromKeyValueStrings(c.Cookies)
	}
	return c.cookieMap
}

func createMapFromKeyValueStrings(kvStrings []string) map[string]string {
	newMap := make(map[string]string)
	for _, kv := range kvStrings {
		parts := strings.Split(kv, "=")
		newMap[parts[0]] = parts[1]
	}
	return newMap
}
