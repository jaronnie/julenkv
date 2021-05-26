package main

import (
	"fmt"
	julenkv "github.com/jaronnie/julenkv/client/go-julenkv"
)

func main() {
	conn, _ := julenkv.Connect("0.0.0.0:5200")
	reply, _ := julenkv.Do(conn, "get", "hello")
	if _, ok := reply.(string); ok {
		value := reply.(string)
		fmt.Println(value)
	}
}
