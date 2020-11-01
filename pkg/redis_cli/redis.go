package redis_cli

import (
	"log"
	"time"

	"github.com/go-redis/redis/v7"
)

type Rcli struct {
	*redis.Client
}

type RedInterface interface {
	Set(key string, value string) error
	Get(key string) (string, error)
}

func NewClient(addr string) (*Rcli, error) {

	client := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     "",
		DB:           0,
		PoolSize:     1000,
		MinIdleConns: 200,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		log.Printf("Redis ping returned error:%v")
		return nil, err
	}
	log.Printf("Ping:%v error: %v", pong, err)
	return &Rcli{client}, nil
}

func (rcli *Rcli) Set(key string, value string) error {
	err := rcli.Client.Set(key, value, 30*time.Minute).Err()
	if err != nil {
		log.Printf("ErrorValue in setting Key:%v", err)
	}
	return err
}

func (rcli *Rcli) Get(key string) (string, error) {
	val, err := rcli.Client.Get(key).Result()

	if err == redis.Nil {
		log.Printf("Redis key does not exist:%v", key)
		return "", err
	}

	if err != nil {
		log.Printf("ErrorValue occured getting redis key:%v error:%v", key, err)
		return "", err
	}

	return val, err
}
