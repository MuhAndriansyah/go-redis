package main

import (
	"log"

	"github.com/MuhAndriansyah/go-redis-crud/task"
	"github.com/hibiken/asynq"
)

func main() {
	srv := asynq.NewServer(asynq.RedisClientOpt{Addr: "localhost:6379"},
		asynq.Config{Concurrency: 10})

	mux := asynq.NewServeMux()
	mux.HandleFunc(task.TypeWelcomeEmail, task.HandleWelcomeEmailTask)
	mux.HandleFunc(task.TypeReminderEmail, task.HandleReminderEmailTask)

	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}
