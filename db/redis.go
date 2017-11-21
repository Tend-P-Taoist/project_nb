package db

import (
	"github.com/garyburd/redigo/redis"
)

var Redis redis.Conn

func init() {
	c,err := redis.Dial("tcp","127.0.0.1:6379")

	if err != nil {
		panic(err.Error())
	}
	Redis = c
}