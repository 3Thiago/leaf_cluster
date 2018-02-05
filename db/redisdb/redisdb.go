package redisdb

import (
	"github.com/garyburd/redigo/redis"
)

type Session struct {
	redis.Conn
}