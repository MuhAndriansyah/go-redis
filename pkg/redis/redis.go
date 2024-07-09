package redis

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type Client struct {
	*redis.Client
}

var (
	once   sync.Once
	client *Client
)

func InitRedis() error {
	var err error
	once.Do(func() {
		// redisHost := os.Getenv("REDIS_HOST")
		// redisPort := os.Getenv("REDIS_HOST")
		// redisPassword := os.Getenv("REDIS_PASSWORD")
		redisHost := "localhost"
		redisPort := "6379"
		redisPassword := ""

		if redisHost == "" || redisPort == "" {
			err = fmt.Errorf("REDIS HOST and PORT must be set")
		}

		conn := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
			Password: redisPassword,
			DB:       0,
			PoolSize: 10,
		})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()

		_, err = conn.Ping(ctx).Result()

		if err != nil {
			err = fmt.Errorf("failed to connect to Redis: %v", err)
			return
		}

		client = &Client{conn}
		logrus.Info("Successfully connected to Redis")

	})

	return err
}

func GetRedis() *Client {
	if client == nil {
		if err := InitRedis(); err != nil {
			logrus.Errorf("Failed to initialize Redis: %v", err)
		}
	}

	return client
}
