package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:5200")
	if err != nil {
		log.Fatal(err)
	}
	conn.Do("Set", "hello", "world")
	reply, _ := conn.Do("Get", "hello")
	fmt.Println(reply)
}
