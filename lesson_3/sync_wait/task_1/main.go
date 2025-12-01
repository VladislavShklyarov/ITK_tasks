package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	cnt := 100
	for i := 0; i < cnt; i++ {
		i := i
		wg.Go(func() {
			fmt.Println(i)
		})
	}
	wg.Wait()
}
