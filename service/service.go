package service

import (
	"law/conf"
	"time"

	"github.com/dgraph-io/ristretto"
	"github.com/eko/gocache/v3/cache"
	"github.com/eko/gocache/v3/store"
)

var CacheManager *cache.Cache[any]

func init() {
	ristrettoCache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: conf.App.Ristretto.NumCounters,
		MaxCost:     conf.App.Ristretto.MaxCost,
		BufferItems: conf.App.Ristretto.BufferItems,
	})
	if err != nil {
		panic(err)
	}
	ristrettoStore := store.NewRistretto(ristrettoCache, store.WithExpiration(60*60*time.Second))
	CacheManager = cache.New[any](ristrettoStore)
}
