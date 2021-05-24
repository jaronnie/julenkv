package cmd

import (
	"errors"
	"fmt"
	"github.com/jaronnie/julenkv"
	"github.com/tidwall/redcon"
)

// ErrSyntaxIncorrect incorrect err
var ErrSyntaxIncorrect = errors.New("syntax err")
var okResult = redcon.SimpleString("OK")

func newWrongNumOfArgsError(cmd string) error {
	return fmt.Errorf("wrong number of arguments for '%s' command", cmd)
}

func set(db *julenkv.JulenKv, args []string) (res interface{}, err error) {
	if len(args) != 2 {
		err = newWrongNumOfArgsError("set")
		return
	}

	key, value := args[0], args[1]
	if err = db.Set([]byte(key), []byte(value)); err == nil {
		res = okResult
	}
	return
}

func get(db *julenkv.JulenKv, args []string) (res interface{}, err error) {
	if len(args) != 1 {
		err = newWrongNumOfArgsError("get")
		return
	}
	key := args[0]
	var val []byte
	if val, err = db.Get([]byte(key)); err == nil {
		res = string(val)
	}
	return
}

func init() {
	addExecCommand("set", set)
	addExecCommand("get", get)
}
