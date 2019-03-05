package cache

import "github.com/garyburd/redigo/redis"

var RedisClient redis.Conn

func SetInit(c redis.Conn){
	RedisClient = c
}

