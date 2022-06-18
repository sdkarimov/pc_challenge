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
	Storage        map[string]interface{}
}

func NewCache(TTL , GCTimer int) *Cache {
	c := Cache{
		TTLSeconds:     TTL,
		GCTimerSeconds: GCTimer,
		mux:            sync.RWMutex{},
		Storage:        make(map[string]interface{})}
	c.RunGC()
	return &c
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mux.RLock()
	res, ok := c.Storage[key]
	c.mux.RUnlock()
	return res, ok
}

func (c *Cache) Set(key string, val interface{}) {
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
				if v.(core.CacheVal).CreateDate+int64(c.TTLSeconds) < now {
					delete(c.Storage, k)
				}
			}
			c.mux.Unlock()
		}
	}(c)
}
