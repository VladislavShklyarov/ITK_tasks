package main

import (
	"fmt"
	"sync"
	"time"
)

const TableAmount int = 5

// В целом могу прдеставить, что по похожему сценарию в проде могут добавлять заказы в очередь для курьеров.
func main() {
	wg := sync.WaitGroup{}

	restaurant := NewRestaurant()

	wg.Add(2)

	go func() {
		defer wg.Done()
		for visitor := range 10 {
			err := restaurant.OccupyTable(visitor)
			if err != nil {
				fmt.Printf("ошибка продюсера: %s", err)
			}
			fmt.Println("Добавили гостя", visitor+1)
		}
		restaurant.Close()
	}()

	go func() {
		defer wg.Done()
		for {
			table, ok := restaurant.ReleaseTable()
			if !ok {
				fmt.Println("ресторан закрыт, конец работы")
				return
			}
			fmt.Println("Ушел гость", table)
		}
	}()

	wg.Wait()

}

type Restaurant struct {
	closed bool
	cond   *sync.Cond
	tables []any
}

func NewRestaurant() *Restaurant {
	return &Restaurant{
		closed: false,
		cond:   sync.NewCond(&sync.Mutex{}),
		tables: make([]any, 0, TableAmount),
	}
}

func (r *Restaurant) OccupyTable(table int) error {
	r.cond.L.Lock()
	defer r.cond.L.Unlock()

	for len(r.tables) == TableAmount && !r.closed {
		fmt.Println("Ресторан переполнен, ожидайте")
		r.cond.Wait()
	}
	if r.closed {
		return fmt.Errorf("ресторан закрыт")
	}
	r.tables = append(r.tables, table)
	r.cond.Signal()
	return nil
}

func (r *Restaurant) ReleaseTable() (any, bool) {
	r.cond.L.Lock()

	for len(r.tables) == 0 && !r.closed {
		r.cond.Wait()
	}

	if len(r.tables) == 0 && r.closed {
		return nil, false
	}
	task := r.tables[0]
	r.tables = r.tables[1:]
	r.cond.Signal()
	r.cond.L.Unlock()
	time.Sleep(500 * time.Millisecond) // имитируем что консьюмеру нужно время на обработку
	return task, true
}

func (r *Restaurant) Close() {
	r.cond.L.Lock()
	r.closed = true
	r.cond.Broadcast()
	r.cond.L.Unlock()
}
