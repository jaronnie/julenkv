package main

import (
	"github.com/jaronnie/julenkv"
	"github.com/jaronnie/julenkv/cmd"
)

func main() {
	server, _ := cmd.NewServer(julenkv.DefaultConfig())
	server.Listen(julenkv.DefaultAddr)
}
