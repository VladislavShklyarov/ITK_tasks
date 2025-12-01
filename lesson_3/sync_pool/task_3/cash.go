package main

import (
	"encoding/json"
	"sync"
	sp "task_3/syncPool"
	"time"
)

type Cache struct {
	vault map[string]item
	mu    sync.RWMutex
	ttl   time.Duration
}

func NewObjectCache(ttl time.Duration) *Cache {
	c := &Cache{
		vault: make(map[string]item),
		mu:    sync.RWMutex{},
		ttl:   ttl,
	}
	c.startExpirationCheck()
	return c
}

type item struct {
	value      any
	expiration time.Time
}

func (c *Cache) Set(key string, value interface{}) {
	newItem := item{
		value:      value,
		expiration: time.Now().Add(c.ttl),
	}
	c.mu.Lock()
	c.vault[key] = newItem
	c.mu.Unlock()
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	if it, ok := c.vault[key]; ok {
		c.mu.RUnlock()
		return it.value, true
	}
	c.mu.RUnlock()
	return nil, false
}

func (c *Cache) startExpirationCheck() {
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for range ticker.C {
			now := time.Now()
			c.mu.Lock()
			for key, it := range c.vault {
				if now.After(it.expiration) {
					delete(c.vault, key)
				}
			}
			c.mu.Unlock()
		}
	}()
}

func (c *Cache) ToJSON() ([]byte, error) {
	c.mu.RLock()

	buf := sp.GetBuffer()

	data := make(map[string]any, len(c.vault))
	for key, it := range c.vault {
		data[key] = it.value
	}

	enc := json.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}

	out := make([]byte, buf.Len())
	copy(out, buf.Bytes())

	c.mu.RUnlock()
	sp.PutBuffer(buf)
	return out, nil
}
