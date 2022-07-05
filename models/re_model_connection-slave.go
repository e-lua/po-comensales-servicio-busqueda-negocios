package models

import (
	"log"
	"math/rand"
	"sync"

	"github.com/gomodule/redigo/redis"
)

type RedisDB_Slave struct {
	redis.Conn
}

var RedisCN_Slave = GetConn()

var (
	once_Slave sync.Once
	p_Slave    *redis.Pool
)

func GetConn_Slave() *redis.Pool {
	once_Slave.Do(func() {
		p_Slave = &redis.Pool{
			Dial: func() (redis.Conn, error) {

				var conn redis.Conn
				var err error
				random := rand.Intn(4)
				if random%2 == 0 {
					conn, err = redis.Dial("tcp", "redis-slave:6379")
				} else {
					conn, err = redis.Dial("tcp", "redis-slave-2:6379")
				}

				if err != nil {
					log.Fatal("ERROR: No se puede conectar con Redis")
				}
				return conn, err
			},
		}
	})

	return p_Slave
}
