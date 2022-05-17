package service

import (
	"law/conf"

	"github.com/dgraph-io/ristretto"
	"github.com/eko/gocache/v2/cache"
	"github.com/eko/gocache/v2/store"
)

var CacheManager cache.CacheInterface

func InitCacheManager() {
	ristrettoCache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: conf.App.Ristretto.NumCounters,
		MaxCost:     conf.App.Ristretto.MaxCost,
		BufferItems: conf.App.Ristretto.BufferItems,
	})
	if err != nil {
		panic(err)
	}
	ristrettoStore := store.NewRistretto(ristrettoCache, nil)
	CacheManager = cache.New(ristrettoStore)
}
