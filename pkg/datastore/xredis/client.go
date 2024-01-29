package xredis

import "github.com/go-redis/redis/v8"

type XDB struct {
	*redis.Client
}
