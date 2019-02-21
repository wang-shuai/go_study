package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/garyburd/redigo/redis"
	"time"
	"os"
	"os/signal"
	"syscall"
)

type RedisConfig struct {
	Address     string `toml:"address"`
	MaxIdle     int `toml:"maxidle"`
	IdleTimeout int `toml:"idletimeout"`
}

//func InitRedis() (RedisConfig,error){
//	var redisConfig RedisConfig
//
//	if _, err := toml.DecodeFile("./config/redisconfig.toml", &redisConfig); err != nil {
//		fmt.Println(err)
//		return redisConfig,err
//	}
//	return redisConfig,nil
//}

var (
	RdsConfig RedisConfig
	Pool *redis.Pool
)

func init() {
	if _, err := toml.DecodeFile("./config/redisconfig.toml", &RdsConfig); err != nil {
		fmt.Println(err)
	}

	redisHost := RdsConfig.Address
	Pool = newPool(redisHost)
	cls()
}

func newPool(server string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     RdsConfig.MaxIdle,
		IdleTimeout: time.Duration(RdsConfig.IdleTimeout) * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func cls() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)

	go func() {
		<-c
		Pool.Close()
		os.Exit(0)
	}()
}
