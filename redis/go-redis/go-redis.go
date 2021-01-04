package go_redis

import (
	"fmt"
	goredis "github.com/go-redis/redis"
	"time"
)

func New(addr, pwd string) (*goredis.Client, error) {
	r := goredis.NewClient(&goredis.Options{
		Addr:         addr,
		Password:     pwd,
		DialTimeout:  time.Second * 5,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		PoolSize:     10000,
	})

	pong, err := r.Ping().Result()
	if err != nil {
		return nil, err
	}
	if pong == "pong" {
		return nil, fmt.Errorf("redis ping result: " + pong + ", addr=" + addr)
	}
	return r, nil
}
