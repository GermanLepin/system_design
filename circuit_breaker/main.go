package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	urls := []string{
		// bad requests
		"https://yeeeeaaaahhhsite.com",
		"https://failedrequest.com",

		// good requests
		"https://github.com",
		"https://amazon.com",

		// bad requests
		"https://comoooon.com",
		"https://veryveryverymuch.com",

		// good requests
		"https://linkedin.com",
		"https://netflix.com",
	}

	suceededRequests := make(chan *http.Response)
	failedRequests := make(chan error)

	circuitBreaker := newCircuitBreaker(3, 10)

	circuitBreaker.wg.Add(2)
	go circuitBreaker.makeRequests(urls, suceededRequests, failedRequests)
	go circuitBreaker.readResponse(suceededRequests, failedRequests)
	circuitBreaker.wg.Wait()
}

func (cb *circuitBreaker) makeRequests(
	urls []string,
	suceededRequests chan *http.Response,
	failedRequests chan error,
) {
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			response, err := http.Get(url)
			if err != nil {
				// circuit breaker implementation
				cb.makeCircuitBreakerRequest(url, suceededRequests, failedRequests)
			} else {
				suceededRequests <- response
			}
		}(url)
	}
	wg.Wait()

	close(failedRequests)
	close(suceededRequests)
	cb.wg.Done()
}

func (cb *circuitBreaker) makeCircuitBreakerRequest(
	url string,
	suceededRequests chan *http.Response,
	failedRequests chan error,
) {
	var wg sync.WaitGroup

	for i := 0; i < cb.maxRequest; i++ {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			response, err := http.Get(url)
			if err == nil {
				suceededRequests <- response
			}
		}(url)
	}
	wg.Wait()

	time.Sleep(time.Duration(cb.delaySeconds) * time.Second)
	response, err := http.Get(url)
	if err != nil {
		failedRequests <- err
	} else {
		suceededRequests <- response
	}
}

func (cb *circuitBreaker) readResponse(
	suceededRequests chan *http.Response,
	failedRequests chan error,
) {
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		for value := range suceededRequests {
			fmt.Println(value)
		}
		wg.Done()
	}()

	go func() {
		for value := range failedRequests {
			fmt.Printf("client: could not create request: %s\n", value)
		}
		wg.Done()
	}()
	wg.Wait()

	cb.wg.Done()
}

type circuitBreaker struct {
	maxRequest   int
	delaySeconds int

	wg sync.WaitGroup
}

func newCircuitBreaker(
	maxRequest int,
	delaySeconds int,
) *circuitBreaker {
	return &circuitBreaker{
		maxRequest:   maxRequest,
		delaySeconds: delaySeconds,
	}
}
