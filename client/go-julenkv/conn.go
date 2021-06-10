package julenkv

import (
	"net"
	"strings"

	"github.com/jaronnie/julenkv/client/go-julenkv/protocol"
)

type Conn struct {
	conn net.Conn
}

func Connect(addr string) (*Conn, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &Conn{conn: conn}, nil
}

func (c *Conn) Do(args ...string) (string, error) {
	var request []string
	for _, v := range args {
		request = append(request, v)
	}

	reqCmd := protocol.GetRequest(request)
	_, err := c.conn.Write(reqCmd)
	if err != nil {
		return "", err
	}

	command := make([]byte, 1024)
	n, err := c.conn.Read(command)
	if err != nil {
		return "", err
	}

	_, err = protocol.GetReply(command[:n])
	if err != nil {
		return "", err
	}

	reply := string(command[:n])
	replys, err := SwitchReply(reply, n)
	return replys, err

}

func SwitchReply(reply string, n int) (string, error) {
	switch string(reply[0]) {
	case "+":
		return TrimRightSpace(reply[1:n]), nil
	case "$":
		return TrimRightSpace(reply[2:n]), nil
	case "-":
		return TrimRightSpace(reply[1:n]), nil
	}
	return "", nil
}

func TrimRightSpace(s string) string {
	return strings.TrimSpace(s)
}


