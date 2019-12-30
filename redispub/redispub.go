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

	cdata := color{"red"}
	jdata, _ := json.Marshal(cdata)

	for i := 0; i < 10; i++ {
		response := client.Publish("color", jdata)
		log.Printf("Send color %v.", cdata.Name)
		if err := response.Err(); err != nil {
			log.Print("Publish error", err)
		}
	}

}
