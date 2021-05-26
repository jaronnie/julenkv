package main

import (
	"github.com/jaronnie/julenkv/cmd/client/protocol"
	"log"
	"net"
	"os"
)

func main() {
	args := os.Args[1:]
	conn, err := Conn("localhost:5200")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	reqCmd := protocol.GetRequest(args)
	_, err = conn.Write(reqCmd)
	if err != nil {
		log.Fatalf("Conn Write err: %v", err)
	}

	command := make([]byte, 1024)
	n, err := conn.Read(command)
	if err != nil {
		log.Fatalf("Conn Read err: %v", err)
	}

	reply, err := protocol.GetReply(command[:n])
	if err != nil {
		log.Fatalf("protocol.GetReply err: %v", err)
	}

	log.Printf("Reply: %v", reply)
	log.Printf("Command: %v", string(command[:n]))

}

func Conn(addr string) (net.Conn, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
