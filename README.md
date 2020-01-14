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