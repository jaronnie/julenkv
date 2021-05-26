package julenkv

import (
	"github.com/jaronnie/julenkv/client/go-julenkv/protocol"
	"net"
)

func Connect(addr string) (net.Conn, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func Do(conn net.Conn, args ...string) (interface{}, error){
	var request []string
	for _, v := range args {
		request = append(request, v)
	}

	reqCmd := protocol.GetRequest(request)
	_, err := conn.Write(reqCmd)
	if err != nil {
		return nil, err
	}

	command := make([]byte, 1024)
	n, err := conn.Read(command)
	if err != nil {
		return nil, err
	}

	_, err = protocol.GetReply(command[:n])
	if err != nil {
		return nil, err
	}

	// bug, fix later...
	return string(command[2:n]), nil
}
