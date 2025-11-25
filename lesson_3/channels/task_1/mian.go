package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Storage struct {
	stored map[int]struct{}
	mu     sync.Mutex
}

func NewStorage() *Storage {
	return &Storage{
		stored: make(map[int]struct{}),
		mu:     sync.Mutex{},
	}
}

func main() {
	storage := NewStorage()
	capacity := 1000
	doubles := make([]int, 0, capacity)
	for i := 0; i < capacity; i++ {
		doubles = append(doubles, rand.Intn(10))
	}
	uniqueIDs := make(chan int, capacity)

	wg := sync.WaitGroup{}
	for i := 0; i < capacity; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			storage.mu.Lock() // Закрываем
			_, ok := storage.stored[doubles[i]]
			if !ok {
				storage.stored[doubles[i]] = struct{}{}
				uniqueIDs <- doubles[i]
			}
			storage.mu.Unlock() // Открываем
		}(i)

	}
	go func() { // горутина для вычитывания
		wg.Wait()
		close(uniqueIDs) // закрываем канал когда все вычитали
	}()

	for val := range uniqueIDs {
		fmt.Println(val)
	}
	fmt.Println(uniqueIDs)
}

// 1 У нас конкурентная запись в мапу. Я бы создал структуру с мьютексом.
// 2 Дождаться выполнения горутин и закрыть канал.
// 3 Не понимаю, зачем мы сначала создаем слайс, кладем в него 1000 элементов, а затем проверяем наличие уникальных
// элементов в мапе, передавая ей в качестве ключа значение, полученное по индексу в слайсе... но ладно, видимо такая
// задумка.
// 4. Как вариант, можно реализовать два лока с повтороной проверкой: RLock на чтение и обычный Lock на запись. Вроде такой
// паттерн называется double checked locking
