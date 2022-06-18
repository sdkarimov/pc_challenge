package storage

import "github.com/sdkarimov/pc_challenge/core"

type Storage interface {
	Get(string) (core.CacheVal, bool)
	Set(string, core.CacheVal)
	RunGC()
}
