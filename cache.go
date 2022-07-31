package cache

import (
	"errors"
	"sync"
	"time"
)

const DefaultDuration = time.Second * 10

type Cache interface {
	Set(key string, value interface{}, durations ...time.Duration)
	Get(key string) (interface{}, error)
	Delete(key string) error
}

type MapCache struct {
	cache    map[string]interface{}
	duration time.Time
	mu       sync.Mutex
}

func New() *MapCache {
	return &MapCache{
		cache:    make(map[string]interface{}),
		duration: time.Now().Add(time.Minute),
		mu:       sync.Mutex{},
	}
}
func (c *MapCache) Set(key string, value interface{}, durations ...time.Duration) {
	duration := DefaultDuration
	// to check if argument is passed
	if len(durations) == 1 {
		duration = durations[0]
	}
	c.mu.Lock()
	c.cache[key] = value
	c.duration = time.Now().Add(duration)
	c.mu.Unlock()
}

func (c *MapCache) Get(key string) (interface{}, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if item, ok := c.cache[key]; ok {
		if time.Now().Before(c.duration) {
			return item, nil
		}
		delete(c.cache, key)
	}
	return nil, errors.New("not found or expired")
}

func (c *MapCache) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.cache[key]; !ok {
		return errors.New("cannot delete, key not found or expired")
	}
	delete(c.cache, key)
	return nil
}
