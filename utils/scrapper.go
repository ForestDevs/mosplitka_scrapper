package utils

import (
	"math/rand"
	"net"
	"net/http"
	"time"

	"github.com/gocolly/colly/v2"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const Domain = "https://mosplitka.ru"

func randomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func OnRequest(c *colly.Collector) {
	c.OnRequest(func(r *colly.Request) {
		// log.Printf("Visit: %s", r.URL)
		r.Headers.Set("User-Agent", randomString())
	})
}

func NewCollector() *colly.Collector {
	c := colly.NewCollector(
		colly.UserAgent("xy"),
		colly.AllowURLRevisit(),
	)

	c.WithTransport(&http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	})

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*httpbin.*",
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	return c
}
