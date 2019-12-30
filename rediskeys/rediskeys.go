package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "nuc:6379",
		Password: "",
		DB:       0,
	})

	startkey := "hello"
	startkey2 := "test"
	startvalue := "world"
	err := client.Set(startkey, startvalue, 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(startkey).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(startkey, val)

	val2, err := client.Get(startkey2).Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println(startkey2, val2)
	}
}
