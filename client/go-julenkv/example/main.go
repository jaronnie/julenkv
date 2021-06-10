package main

import (
	"fmt"
	"github.com/jaronnie/julenkv/client/go-julenkv"
)

func main() {
	conn, _ := julenkv.Connect("0.0.0.0:5200")
	result, _ := conn.Do("set", "nj", "jay")
	fmt.Println(result)
	result, _ = conn.Do("get", "nj")
	fmt.Println(result)
}
