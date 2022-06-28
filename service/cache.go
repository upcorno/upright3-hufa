package service

import (
	"context"
	"fmt"
	"law/conf"

	"github.com/go-redis/redis/v8"
	zlog "github.com/rs/zerolog/log"
)

var Rdb *redis.Client

func init() {
	if conf.TestMode {
		return
	}
	addr := fmt.Sprintf("%s:%d", conf.App.Rdb.RdbHost, conf.App.Rdb.RdbPort)
	Rdb = redis.NewClient(&redis.Options{
		Addr:       addr,
		Password:   conf.App.Rdb.RdbPasswd,
		DB:         conf.App.Rdb.DbIndex,
		MaxRetries: conf.App.Rdb.MaxRetries,
	})
	_, err := Rdb.Ping(context.Background()).Result()
	if err != nil {
		zlog.Fatal().Err(err).Msgf("redis ping失败.err.")
	}
	zlog.Info().Msgf("成功连接redis")
}
