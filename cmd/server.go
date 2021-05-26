package cmd

import (
	"fmt"
	"github.com/jaronnie/julenkv/config"
	"log"
	"strings"

	"github.com/jaronnie/julenkv"
	"github.com/tidwall/redcon"
)

type ExecCmdFunc func(*julenkv.JulenKv, []string) (interface{}, error)

var ExecCmd = make(map[string]ExecCmdFunc)

func addExecCommand(cmd string, cmdFunc ExecCmdFunc) {
	ExecCmd[strings.ToLower(cmd)] = cmdFunc
}

type Server struct {
	server *redcon.Server
	db     *julenkv.JulenKv
	closed bool
}

func NewServer(config config.Config) (*Server, error) {
	db, err := julenkv.Open(config)
	if err != nil {
		return nil, err
	}
	return &Server{db: db}, nil
}

func (s *Server) Listen(addr string) {
	svr := redcon.NewServerNetwork("tcp", addr,
		func(conn redcon.Conn, cmd redcon.Command) {
			s.handleCmd(conn, cmd)
		},
		func(conn redcon.Conn) bool {
			return true
		},
		func(conn redcon.Conn, err error) {
		},
	)
	s.server = svr
	if err := svr.ListenAndServe(); err != nil {
		log.Printf("listen and serve ocuurs error: %+v", err)
	}
}

// Stop stops the server
func (s *Server) Stop() {
	if s.closed {
		return
	}
	s.closed = true
	if err := s.server.Close(); err != nil {
		log.Printf("close redcon err: %+v\n", err)
	}
	if err := s.db.Close(); err != nil {
		log.Printf("close julenkv err: %+v\n", err)
	}
}

func (s *Server) handleCmd(conn redcon.Conn, cmd redcon.Command) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("panic when handle the cmd: %+v", r)
		}
	}()

	command := strings.ToLower(string(cmd.Args[0]))
	exec, exist := ExecCmd[command]
	if !exist {
		conn.WriteError(fmt.Sprintf("ERR unknown command '%s'", command))
		return
	}
	args := make([]string, 0, len(cmd.Args)-1)
	for i, bytes := range cmd.Args {
		if i == 0 {
			continue
		}
		args = append(args, string(bytes))
	}
	reply, err := exec(s.db, args)
	if err != nil {
		conn.WriteError(err.Error())
		return
	}
	conn.WriteAny(reply)
}
