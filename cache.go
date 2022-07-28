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
	sync.RWMutex
}

func New() *MapCache {
	return &MapCache{
		cache: make(map[string]interface{}),
	}
}
func (c *MapCache) Set(key string, value interface{}) {
	c.Lock()
	c.cache[key] = value
	c.Unlock()
}

func (c *MapCache) Get(key string) (interface{}, error) {
	c.RLock()
	defer c.RUnlock()
	if _, ok := c.cache[key]; !ok {
		return nil, errors.New("not found")
	}
	return c.cache[key], nil
}

func (c *MapCache) Delete(key string) error {
	c.Lock()
	defer c.Unlock()
	if _, ok := c.cache[key]; !ok {
		return errors.New("cannot delete, key not found")
	}
	delete(c.cache, key)
	return nil
}
