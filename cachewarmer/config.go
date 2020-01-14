package cachewarmer

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

type Config struct {
	Type              string
	Path              string
	Cookies           []string
	Headers           []string
	HttpClientTimeout time.Duration
}

// Checks if the configuration is valid, returns error if invalid.
func (c *Config) Validate() error {
	if c.Type != "list" && c.Type != "sitemap" {
		return errors.New("type is invalid, should be either 'list' or 'sitemap'")
	}

	// if it doesn't start with http it's probably a file path, check if that file exists...
	if !strings.HasPrefix(c.Path, "http") {
		_, err := os.Stat(c.Path)
		if os.IsNotExist(err) {
			return errors.New(fmt.Sprintf("file %s does not exist", c.Path))
		}
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
	for _, keyValuePair := range keyValueSet {
		if matched, _ := regexp.MatchString(`.*=.*`, keyValuePair); !matched {
			return errors.New(fmt.Sprintf(
				"%s does not match pattern %s_name=%s_value for: %s",
				name,
				name,
				name,
				keyValuePair,
			))
		}
	}
}
