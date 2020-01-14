package cachewarmer

import (
	"log"
	"net/http"
	"sync"
)

type Crawler struct {
	Client          http.Client
	NumberOfWorkers int
	Logger          *log.Logger
}

func (c *Crawler) Crawl(requests []*http.Request) {
	requestsNum := len(requests)

	queue := make(chan *http.Request, requestsNum)
	for _, req := range requests {
		queue <- req
	}

	wg := sync.WaitGroup{}
	wg.Add(requestsNum)

	for i := 0; i < c.NumberOfWorkers; i++ {
		go func() {
			for req := range queue {
				c.crawlRequest(req)
				wg.Done()
			}
		}()
	}

	wg.Wait()
	close(queue)
}

func (c *Crawler) crawlRequest(req *http.Request) {
	res, err := c.Client.Do(req)
	if err != nil {
		c.Logger.Printf("error with url: %s\n%e\n", req.URL, err)
		return
	}

	c.Logger.Println(req.URL, res.StatusCode)
}
