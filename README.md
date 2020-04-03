# Discontinued use [github.com/atomicptr/crab](https://github.com/atomicptr/crab) instead

Since I started using this tool for much more than cache warming (e.g. testing if urls are still working),
I've decided to discontinue this tool in favor of a new one called [crab](https://github.com/atomicptr/crab).

Crab has every feature this one has and more.

## Basic Usage with Docker

```bash
$ docker run --rm atomicptr/crab crawl:sitemap https://domain.com/sitemap.xml
```

# cache-warmer

Simple cache warmer which reads a sitemap.xml and requests all URLs within it.

## Usage

```
Usage: cache-warmer [options] [arguments]

OPTIONS
    --provider/$CW_PROVIDER  <string>  (required)
    How should the tool query requests?

    --path/$CW_PATH  <string>  (required)
    Path to the URL list or Path/URL to the sitemap.xml

    --cookies/$CW_COOKIES  <string>,[string...]
    Cookies to add to the request

    --headers/$CW_HEADERS  <string>,[string...]
    Headers to add to the request

    --prefix-url/$CW_PREFIX_URL  <string>
    Prefix an URL or replace the URL altogether

    --http-client-timeout/$CW_HTTP_CLIENT_TIMEOUT  <duration>  (default: 30s)

    --number-of-workers/$CW_NUMBER_OF_WORKERS      <int>       (default: 32)

    --help/-h
    display this help message

```

## Basic usage with Docker

```bash
$ docker run --rm \
    -e CW_PROVIDER=sitemap \
    -e CW_PATH="https://example.com/sitemap.xml" \
    atomicptr/cache-warmer
```

## Basic usage with executable

```bash
$ ./cache-warmer --provider=sitemap --path=https://example.com/sitemap.xml
```

## License

MIT