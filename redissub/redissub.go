package main

import (
	"encoding/json"
	"log"

	"github.com/go-redis/redis"
)

type color struct {
	Name string
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "nuc:6379",
		Password: "",
		DB:       0,
	})

	pubsub := client.Subscribe("color")
	defer pubsub.Close()

	_, err := pubsub.Receive()
	if err != nil {
		panic(err)
	}

	log.Printf("Listening for colors")
	channel := pubsub.Channel()

	for msg := range channel {
		var responsevalue color

		json.Unmarshal([]byte(msg.Payload), &responsevalue)
		log.Printf("Recieved color %v.", responsevalue.Name)
	}
}
