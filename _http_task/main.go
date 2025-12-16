package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	limiter := NewRateLimier(500)

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Go(func() {
			err := makeRequest(limiter, "https://example.com", "{{}}")
			if err != nil {
				fmt.Println("request", i, "error:", err)
			} else {
				fmt.Println("request", i, "done")
			}
		})
	}

	wg.Wait()
}

type RateLimiter struct {
	tokens chan struct{}
}

func NewRateLimier(rps int) *RateLimiter {
	rl := &RateLimiter{tokens: make(chan struct{}, rps)} // создаем буферизованный канал

	ticker := time.NewTicker(time.Second / time.Duration(rps))

	go func() {
		for range ticker.C { // вычитываем из канала тикера
			select {
			case rl.tokens <- struct{}{}:
			default:
			}
		}
	}()

	return rl
}

func makeRequest(rl *RateLimiter, url, data string) error {
	<-rl.tokens

	fmt.Printf("requesting %s: %s\n", url, data)
	return nil
}
