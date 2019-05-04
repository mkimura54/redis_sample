package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	red, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
	}
	defer red.Close()

	red.Do("SET", "abc", 123)
	v, err := redis.String(red.Do("GET", "abc"))
	if err != nil {
		panic(err)
	}
	fmt.Println(v)
}
