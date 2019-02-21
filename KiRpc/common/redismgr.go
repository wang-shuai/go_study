package common

import (
	"os"
	"github.com/garyburd/redigo/redis"
	"time"
	"syscall"
	"os/signal"
)

var (
	Pool *redis.Pool
)

func init() {
	Pool = newPool()
	cls()
}

func newPool() *redis.Pool {
	address, _ := CfgMgr.String("redis.address")
	maxIdle, _ := CfgMgr.Int("redis.maxidle")
	idelTimeout, _ := CfgMgr.Int("redis.idletimeout")
	maxActive, _ := CfgMgr.Int("redis.redismaxactive")
	return &redis.Pool{
		MaxIdle:     maxIdle,
		IdleTimeout: time.Duration(idelTimeout) * time.Second,
		MaxActive:   maxActive,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", address)
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
