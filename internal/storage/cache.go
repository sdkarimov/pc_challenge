package storage

import (
	"sync"
	"time"

	"github.com/sdkarimov/pc_challenge/core"
)

type Cache struct {
	TTLSeconds     int
	GCTimerSeconds int
	mux            sync.RWMutex
	Storage        map[string]core.CacheVal
}

func NewCache(ttl int, gctimer int) *Cache {
	c := Cache{
		TTLSeconds:     ttl,
		GCTimerSeconds: gctimer,
		mux:            sync.RWMutex{},
		Storage:        make(map[string]core.CacheVal)}
	c.RunGC()
	return &c
}

func (c *Cache) Get(key string) (core.CacheVal, bool) {
	c.mux.RLock()
	res, ok := c.Storage[key]
	c.mux.RUnlock()
	return res, ok
}

func (c *Cache) Set(key string, val core.CacheVal) {
	c.mux.Lock()
	c.Storage[key] = val
	c.mux.Unlock()
}

func (c *Cache) RunGC() {
	go func(c *Cache) {
		for {
			time.Sleep(time.Duration(c.GCTimerSeconds) * time.Second)
			c.mux.Lock()
			for k, v := range c.Storage {
				now := time.Now().Unix()
				if v.CreateDate+int64(c.TTLSeconds) < now {
					delete(c.Storage, k)
				}
			}
			c.mux.Unlock()
		}
	}(c)
}
