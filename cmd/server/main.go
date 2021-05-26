package main

import (
	"github.com/jaronnie/julenkv/cmd"
	"github.com/jaronnie/julenkv/config"
)

func main() {
	server, _ := cmd.NewServer(config.DefaultConfig())
	server.Listen(config.DefaultAddr)
}
