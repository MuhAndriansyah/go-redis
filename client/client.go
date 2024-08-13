package main

import (
	"log"
	"time"

	"github.com/MuhAndriansyah/go-redis-crud/task"
	"github.com/hibiken/asynq"
)

type EmailTaskPayload struct {
	UserID int
}

//Konsep client pada asynq ini, menerima sebuah task yang nanti nya akan kita
//proses (enqueue) ke redis

func main() {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"})

	t1, err := task.NewWelcomeEmailTask(42)
	if err != nil {
		log.Fatal(err)
	}

	t2, err := task.NewReminderEmailTask(42)
	if err != nil {
		log.Fatal(err)
	}

	info, err := client.Enqueue(t1)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[*] successfully enqueued task: %+v", info)

	info, err = client.Enqueue(t2, asynq.ProcessIn(5*time.Minute))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[*] successfully enqueued task: %+v", info)
}
