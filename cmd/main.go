package main

import (
	"github.com/MuhAndriansyah/go-redis-crud/cmd/api"
	"github.com/MuhAndriansyah/go-redis-crud/internal/book"
	"github.com/MuhAndriansyah/go-redis-crud/pkg/redis"
	"github.com/sirupsen/logrus"
)

func main() {

	logrus.SetLevel(logrus.DebugLevel)

	book.InitializeBooks()
	redis.InitRedis()

	server := api.NewAPIServer(":3000")

	if err := server.Run(); err != nil {
		logrus.Fatalf("failed to start the server %v", err)
		return
	}

	logrus.Info("server stoped")
}
