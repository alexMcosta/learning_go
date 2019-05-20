package main

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	aDur := measureResponseTime(a)
	bDur := measureResponseTime(b)

	if aDur > bDur {
		return b
	}

	return a
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func main() {}
