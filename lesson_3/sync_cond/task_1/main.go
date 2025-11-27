package main

import (
	"fmt"
	"sync"
	"time"
)

const QueueCap int = 5

// нашел вот такую статью по теме https://ubiklab.net/posts/go-sync-cond/
// в целом вроде разобрался, но не понял в чем большой плюс перед решением с буферизованными каналами.
func main() {
	wg := sync.WaitGroup{}

	boundQueue := NewBQ()

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := range 10 {
			err := boundQueue.Put(i)
			if err != nil {
				fmt.Printf("ошибка продюсера: %s", err)
			}
			fmt.Println("Добавили", i)
		}
		boundQueue.ShutDown()
	}()

	go func() {
		defer wg.Done()
		for {
			task, ok := boundQueue.Get()
			if !ok {
				fmt.Println("очередь закрыта, конец работы")
				return
			}
			fmt.Println("Взяли задачу", task)
		}
	}()

	wg.Wait()

}

type BoundQueue struct {
	closed bool
	cond   *sync.Cond
	queue  []any
}

func NewBQ() *BoundQueue {
	return &BoundQueue{
		closed: false,
		cond:   sync.NewCond(&sync.Mutex{}),
		queue:  make([]any, 0, QueueCap),
	}
}

func (bq *BoundQueue) Put(task any) error {
	bq.cond.L.Lock()
	defer bq.cond.L.Unlock()

	for len(bq.queue) == QueueCap && !bq.closed {
		fmt.Println("Очередь переполнена, ожидайте")
		bq.cond.Wait()
	}
	if bq.closed {
		return fmt.Errorf("очередь закрыта")
	}
	bq.queue = append(bq.queue, task)
	bq.cond.Signal()
	return nil
}

func (bq *BoundQueue) Get() (any, bool) {
	bq.cond.L.Lock()

	for len(bq.queue) == 0 && !bq.closed {
		bq.cond.Wait()
	}

	if len(bq.queue) == 0 && bq.closed {
		return nil, false
	}
	task := bq.queue[0]
	bq.queue = bq.queue[1:]
	bq.cond.Signal()
	bq.cond.L.Unlock()
	time.Sleep(500 * time.Millisecond) // имитируем что консьюмеру нужно время на обработку
	return task, true
}

func (bq *BoundQueue) ShutDown() {
	bq.cond.L.Lock()
	bq.closed = true
	bq.cond.Broadcast()
	bq.cond.L.Unlock()
}
