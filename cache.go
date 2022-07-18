package cache

import "errors"

type Cache interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, error)
	Delete(key string)
}

type MapCache struct {
	cache map[string]interface{}
}

func New() *MapCache {
	return &MapCache{
		cache: make(map[string]interface{}),
	}
}
func (c *MapCache) Set(key string, value interface{}) {
	c.cache[key] = value
}

func (c MapCache) Get(key string) (interface{}, error) {
	if _, ok := c.cache[key]; !ok {
		return nil, errors.New("not found")
	}
	return c.cache[key], nil
}

func (c *MapCache) Delete(key string) {
	delete(c.cache, key)
}
