package db

import (
	"crypto/tls"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"sf-go/logs"
	"time"
)

type RDB struct {
	Rdb *redis.Client
}

func NewRedisDB(
	wcfg RedisConfig,
) (*RDB, error) {
	rdb, err := SetRedis(wcfg)
	if err != nil {
		return nil, err
	}
	return rdb, nil
}

func SetRedis(wcfg RedisConfig) (*RDB, error) {

	logs.Logger.Info("init redis start...")
	rdb := &RDB{}

	if wcfg.IsTls {
		rdb.Rdb = redis.NewClient(&redis.Options{
			Addr:     wcfg.Addr,
			Password: wcfg.Password, // no password set
			DB:       wcfg.DB,       // use default DB
			TLSConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		})

	} else {
		rdb.Rdb = redis.NewClient(&redis.Options{
			Addr:     wcfg.Addr,
			Password: wcfg.Password, // no password set
			DB:       wcfg.DB,       // use default DB
		})
	}
	pingResult, err := rdb.Rdb.Ping().Result()
	if err != nil {
		logs.Logger.Fatal("Init redis error :", zap.Error(err))
		return nil, err
	}
	logs.Logger.Info("redis ping results:", zap.String("pingResult", pingResult))
	go func() {
		for {
			time.Sleep(5 * time.Second)
			_, err = rdb.Rdb.Ping().Result()
			if err != nil {
				_ = rdb.Rdb.Close()
				logs.Logger.Info("PING Failure: ", zap.Error(err))
				_, err = SetRedis(wcfg)
				if err != nil {
					logs.Logger.Fatal("Init redis error :", zap.Error(err))
				}
			}
		}
	}()
	return rdb, nil
}
