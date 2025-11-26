package main

import (
	"fmt"
	"sync"
)

func main() {
	cache := NewSafeCache()

	var wg sync.WaitGroup

	// Параллельная запись в кеш
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key_%d", i)
			value := fmt.Sprintf("value_%d", i)
			cache.Set(key, value)
			fmt.Println("Set:", key, value)
		}(i)
	}

	wg.Wait()

	// Параллельное чтение из кеша
	for i := 0; i < 12; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key_%d", i)
			if value, ok := cache.Get(key); ok {
				fmt.Println("Get:", key, value)
			} else {
				fmt.Println("Get:", key, "not found")
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("Все операции завершены.")
}

type SafeCache struct {
	vault map[string]string
	mu    sync.RWMutex
}

func NewSafeCache() *SafeCache {
	return &SafeCache{
		vault: make(map[string]string),
		mu:    sync.RWMutex{},
	}
}

func (sc *SafeCache) Set(key, value string) {
	sc.mu.Lock()
	if _, ok := sc.vault[key]; !ok {
		sc.vault[key] = value
	}
	sc.mu.Unlock()
}

func (sc *SafeCache) Get(key string) (string, bool) {
	sc.mu.RLock()
	defer sc.mu.RUnlock()
	value, ok := sc.vault[key]
	return value, ok
}
