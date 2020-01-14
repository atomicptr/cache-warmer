package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ardanlabs/conf"
	"github.com/pkg/errors"

	"github.com/atomicptr/cache-warmer/cachewarmer"
	_ "github.com/atomicptr/cache-warmer/urlproviders"
)

const ConfNamespace = "CW"

func main() {
	err := run()
	if err != nil {
		log.Printf("error: %s", err)
	}
}

func run() error {
	// logger
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	// configuration

	var config struct {
		Provider          string        `conf:"required,help:How should the tool query requests? Provide either 'list' or 'sitemap'"`
		Path              string        `conf:"required,help:Path to the URL list or Path/URL to the sitemap.xml"`
		Cookies           []string      `conf:"help:Cookies to add to the request"`
		Headers           []string      `conf:"help:Headers to add to the request"`
		PrefixUrl         string        `conf:"help:Prefix an URL or replace the URL altogether"`
		HttpClientTimeout time.Duration `conf:"default:30s"`
		NumberOfWorkers   int           `conf:"default:32"`
	}

	if err := conf.Parse(os.Args[1:], ConfNamespace, &config); err != nil {
		if err == conf.ErrHelpWanted {
			usage, err := conf.Usage(ConfNamespace, &config)
			if err != nil {
				return errors.Wrap(err, "generating usage")
			}

			fmt.Println(usage)
			return nil
		}
		return errors.Wrap(err, "error: parsing config")
	}

	// app starting
	logger.Printf("main: cache-warmer starting...")
	defer logger.Printf("main: Done")

	out, err := conf.String(&config)
	if err != nil {
		return errors.Wrap(err, "generating config for output")
	}
	logger.Printf("main: Config:\n%v\n", out)

	// channel to listen for interrupt or terminate signal from OS
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// init cache warmer
	cacheWarmer, err := cachewarmer.New(
		cachewarmer.Config{
			Provider:          config.Provider,
			Path:              config.Path,
			Cookies:           config.Cookies,
			Headers:           config.Headers,
			PrefixUrl:         config.PrefixUrl,
			HttpClientTimeout: config.HttpClientTimeout,
			NumberOfWorkers:   config.NumberOfWorkers,
		},
		logger,
	)

	if err != nil {
		return errors.Wrap(err, "configuration is invalid")
	}

	// channel to listen for errors coming from the cache warmer
	cacheWarmerErrors := make(chan error, 1)

	go func() {
		cacheWarmerErrors <- cacheWarmer.Run()
	}()

	select {
	case err := <-cacheWarmerErrors:
		return errors.Wrap(err, "cache warmer error")
	case sig := <-shutdown:
		logger.Printf("main: %v shutdown...", sig)
		switch {
		case sig == syscall.SIGSTOP:
			return errors.New("integrity issue caused shutdown")
		}
	}

	return nil
}
