package main

import (
	"sync"
	"time"
)

const workerNums = 1000

var wg sync.WaitGroup

func worker() chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- 42
	}()
	return ch
}
func main() {

	timeStart := time.Now()
	for range workerNums {
		wg.Go(func() {
			<-worker()
		})
	}

	wg.Wait()
	println(int(time.Since(timeStart).Seconds()))
}

// изначально мы создаем 2 воркера, и они выполняются конкурентно. Сначала первый, затем второй. Поэтому изначально
// время выполнения программы 6 секунд.
// выход: можно запустить горутину на каждого воркера. Тогда они будут выполняться паралельно, и время выполнения будет
// равно времени работы одного воркера
