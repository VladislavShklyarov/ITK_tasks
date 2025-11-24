package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

type Cash[T any] struct {
	cash map[string]*Item
	mu   sync.RWMutex
}

func NewCache[T any]() *Cash[T] {
	cash := &Cash[T]{
		cash: make(map[string]*Item),
		mu:   sync.RWMutex{},
	}
	return cash
}

type Item struct {
	Value        any
	CreationTime time.Time
	Ttl          time.Duration
}

func (i *Item) isExpired() (int, bool) {
	remaining := i.Ttl - time.Since(i.CreationTime)
	return int(remaining.Seconds()), remaining < 0
}

func newItem(value any, ttl time.Duration) *Item {
	return &Item{
		Value:        value,
		CreationTime: time.Now(),
		Ttl:          ttl,
	}
}

func (c *Cash[T]) Set(key string, value any, ttl time.Duration) {
	data := newItem(value, ttl)
	c.mu.Lock()
	c.cash[key] = data
	c.mu.Unlock()
}

func (c *Cash[T]) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	element, ok := c.cash[key]
	c.mu.RUnlock()
	if !ok {
		fmt.Print(fmt.Errorf("такого ключа нет\t"))
		return nil, false
	}

	remaining, expired := element.isExpired()
	if expired {
		fmt.Println("время жизни объекта истекло")
		c.Delete(key)
		return element.Value, false
	}
	fmt.Printf("Объекту %s осталось жить %v секунд\n", key, remaining)
	return element.Value, true
}

func (c *Cash[T]) Delete(key string) {
	c.mu.Lock()
	delete(c.cash, key)
	c.mu.Unlock()
}

func (c *Cash[T]) Clear() {
	for key := range c.cash {
		delete(c.cash, key)
	}
}

func (c *Cash[T]) Exists(key string) bool {
	item, ok := c.cash[key]
	if !ok {
		return false
	}
	_, expired := item.isExpired()
	return !expired
}

func (c *Cash[T]) ToJSON() ([]byte, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return json.Marshal(c.cash)
}

func (c *Cash[T]) GetAs(key string) (interface{}, error) {
	c.mu.RLock()
	element, ok := c.cash[key]
	c.mu.RUnlock()

	if !ok {
		return nil, fmt.Errorf("ключ %s не найден", key)
	}

	_, expired := element.isExpired()
	if expired {
		c.Delete(key)
		return nil, fmt.Errorf("время жизни ключа %s истекло", key)
	}

	val := element.Value
	fmt.Printf("Тип %T", val)

	return val, nil
}
