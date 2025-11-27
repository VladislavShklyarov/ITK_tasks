package main

import (
	"fmt"
	"sync"
	"time"
)

const QueueCap = 5

func main() {
	wg := sync.WaitGroup{}
	boundQueue := NewBoundQueue()

	wg.Add(4)
	go func() {
		defer wg.Done()
		for i := range 10 {
			boundQueue.Put(i)
			fmt.Printf("Положили %d\n", i)
		}
		boundQueue.Shutdown()
	}()

	for i := range 3 { // несколько воркеров
		go func() {
			defer wg.Done()
			for {
				task, ok := boundQueue.Get()
				if !ok {
					fmt.Println("очередь закрыта, конец работы")
					return
				}
				fmt.Printf("Воркер %d, взял %d\n", i, task)
			}
		}()
	}

	wg.Wait()
}

type BoundQueue struct {
	queue chan any
}

func NewBoundQueue() *BoundQueue {
	return &BoundQueue{
		queue: make(chan any, QueueCap),
	}
}

func (bq *BoundQueue) Put(task any) {
	bq.queue <- task
}

func (bq *BoundQueue) Get() (any, bool) {
	time.Sleep(500 * time.Millisecond)
	task, ok := <-bq.queue
	return task, ok
}

func (bq *BoundQueue) Shutdown() {
	close(bq.queue)
}
