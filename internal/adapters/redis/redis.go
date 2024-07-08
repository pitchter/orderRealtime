package redis

import (
    "context"
    "github.com/redis/go-redis/v9"
    "log"
)

var ctx = context.Background()
var Rdb *redis.Client
var Nil = redis.Nil

// Init initializes the Redis client and returns it.
func Init() *redis.Client {
    Rdb = redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
    })

    _, err := Rdb.Ping(ctx).Result()
    if err != nil {
        log.Fatalf("Could not connect to Redis: %v", err)
    }

    return Rdb
}


