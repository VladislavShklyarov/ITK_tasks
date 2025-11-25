package main

import (
	"fmt"
	"sync"
)

func mergeChannels(channels ...<-chan int) <-chan int {
	wg := sync.WaitGroup{}

	res := make(chan int)
	for _, ch := range channels {
		wg.Go(func() {
			for el := range ch {
				res <- el
			}
		})
	}
	go func() {
		wg.Wait()
		close(res)
	}()

	return res
}

func main() {
	wg := sync.WaitGroup{}

	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	wg.Go(func() {
		defer close(a)
		for range 10 {
			a <- 1
		}

	})
	wg.Go(func() {
		defer close(b)
		for range 10 {
			b <- 2
		}
	})
	wg.Go(func() {
		defer close(c)
		for range 10 {
			c <- 3
		}
	})

	go func() {
		wg.Wait()
	}()

	res := mergeChannels(a, b, c)
	for el := range res {
		fmt.Println(el)
	}

}
