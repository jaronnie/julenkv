package main

import (
	"github.com/jaronnie/julenkv/config"
	"github.com/jaronnie/julenkv/server"
)

func main() {
	srv, _ := server.NewServer(config.DefaultConfig())
	srv.Listen(config.DefaultAddr)
}
