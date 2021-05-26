package main

import (
	"fmt"
	julenkv "github.com/jaronnie/julenkv/client/go-julenkv"
)

func main() {
	conn, _ := julenkv.Connect("0.0.0.0:5200")
	julenkv.Do(conn, "set", "nj", "jay")
	reply, _ := julenkv.Do(conn, "get", "nj")
	if _, ok := reply.(string); ok {
		value := reply.(string)
		fmt.Println(value)
	}
}
