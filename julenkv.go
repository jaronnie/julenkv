package julenkv

import (
	"github.com/jaronnie/julenkv/config"
	"github.com/jaronnie/julenkv/ds/dict"
)

type JulenKv struct {
	dictIdx *dict.Dict
}

func Open(config config.Config) (*JulenKv, error) {
	return &JulenKv{dict.NewDict()}, nil
}

func (db *JulenKv) Close() error {
	//TODO
	return nil
}
