package cache

import (
	"errors"
	"sync"
)

type Cache interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, error)
	Delete(key string)
}

type MapCache struct {
	cache map[string]interface{}
	mu    sync.RWMutex
}

func New() *MapCache {
	return &MapCache{
		cache: make(map[string]interface{}),
		mu:    sync.RWMutex{},
	}
}
func (c *MapCache) Set(key string, value interface{}) {
	c.mu.Lock()
	c.cache[key] = value
	c.mu.Unlock()
}

func (c *MapCache) Get(key string) (interface{}, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if _, ok := c.cache[key]; !ok {
		return nil, errors.New("not found")
	}
	return c.cache[key], nil
}

func (c *MapCache) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.cache[key]; !ok {
		return errors.New("cannot delete, key not found")
	}
	delete(c.cache, key)
	return nil
}
