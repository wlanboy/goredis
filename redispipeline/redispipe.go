package main

import (
	"fmt"
	"time"

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

	pipe := client.Pipeline()

	fmt.Println("creating pipeline counter")
	incr := pipe.Incr("pipeline_counter")
	incr2 := pipe.Incr("pipeline_counter")
	pipe.Expire("pipeline_counter", time.Hour)

	fmt.Println("excecuting pipeline")
	_, err := pipe.Exec()
	fmt.Println(incr.Val(), err)
	fmt.Println(incr2.Val(), err)

}
