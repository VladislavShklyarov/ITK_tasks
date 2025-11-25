package main

import (
	"fmt"
	"sync"
	"time"
)

// Реплика БД (имитация)
func dbReplica(name string, in <-chan int) {
	for data := range in {
		fmt.Printf("Запись в %s: %d\n", name, data)
		time.Sleep(100 * time.Millisecond) // Имитация задержки записи
	}
	fmt.Printf("Реплика %s закрыта\n", name)
}

func fillInput() <-chan int {
	ch := make(chan int)
	go func() {
		for i := range 10 { // Добавляем данные в input
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func main() {

	var wg sync.WaitGroup

	input := fillInput()    // Канал для входящих данных
	replicas := []chan int{ // Реплики БД (каналы)
		make(chan int),
		make(chan int),
		make(chan int),
	}

	// запускаем параллельно реплики
	for i, replica := range replicas {
		wg.Go(func() {
			dbReplica(fmt.Sprintf("реплика %d", i+1), replica)
		})
	}

	go func() {
		for data := range input {
			for _, replica := range replicas {
				replica <- data // Отправляем данные в реплики
			}
		}

		for _, replica := range replicas {
			close(replica) // Закрываем
		}
	}()

	wg.Wait()

	fmt.Println("Done")
}

// TODO: доработать гарантию записи в реплику
