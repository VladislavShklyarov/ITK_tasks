package main

import (
	"fmt"
	"math/rand"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for range 10 {
			naturals <- rand.Intn(100)
		}
		close(naturals)
	}()

	go func() {
		for num := range naturals {
			squares <- num * num
		}
		close(squares)
	}()

	for el := range squares {
		fmt.Println(el)
	}
}
