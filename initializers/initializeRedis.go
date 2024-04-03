package initializers

import (
	"github.com/go-redis/redis/v8"
	"log"
	"os"
	"strconv"
)

var RedisClient *redis.Client

func InitializeRedis() {

	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		log.Fatal("There was an error reading data for redis")
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:" + os.Getenv("REDIS_PORT"),
		Password: "",
		DB:       db,
	})
}
