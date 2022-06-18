package storage

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/sdkarimov/pc_challenge/core"
	"github.com/stretchr/testify/assert"
)

func TestNewCache(t *testing.T) {
	c := NewCache(3, 1)
	if reflect.TypeOf(c).String() != "*storage.Cache" {
		t.Fatal("NewCache not return Cache obj")
	}
}

func TestCacheUsage(t *testing.T) {
	c := NewCache(3, 1)

	val := core.CacheVal{Value: "testval", CreateDate: time.Now().Unix()}
	c.Set("key", val)
	val1 := core.CacheVal{Value: "testval1", CreateDate: time.Now().Unix()}
	c.Set("key1", val1)
	val2 := core.CacheVal{Value: "testval2", CreateDate: time.Now().Unix()}
	c.Set("key2", val2)

	if v, ok := c.Get("key"); !ok || v.(core.CacheVal).Value != val.Value {
		t.Fatal("Fail to get key = ", val.Value)
	}
	if v, ok := c.Get("key1"); !ok || v.(core.CacheVal).Value != val1.Value {
		t.Fatal("Fail to get key = ", val1.Value)
	}
	if v, ok := c.Get("key2"); !ok || v.(core.CacheVal).Value != val2.Value {
		t.Fatal("Fail to get key = ", val2.Value)
	}

}

func TestCacheGC(t *testing.T) {
	c := NewCache(2, 1)
	val := core.CacheVal{Value: "testval", CreateDate: time.Now().Unix()}
	c.Set("key", val)
	if v, ok := c.Get("key"); !ok || v.(core.CacheVal).Value != val.Value {
		t.Fatal("Fail to get key = ", val.Value)
	}
	fmt.Println("Sleeping 3sec until gc done ")
	time.Sleep(3 * time.Second)
	if v, ok := c.Get("key"); ok {
		assert.Equal(t, v.(core.CacheVal).Value, val.Value)
		t.Fatal("Failed gc worker")
	}

}
