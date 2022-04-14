package tools

import (
	"fmt"
	"sync"
)

type Tool interface {
	Process(ctx ProcessContext) (ProcessContext, error)
}

type Tools struct {
	tools map[string]Tool
	m     sync.Mutex
}

func NewTools() *Tools {
	return &Tools{
		tools: map[string]Tool{},
	}
}

func (t *Tools) Register(name string, tool Tool) {
	t.m.Lock()
	defer t.m.Unlock()
	t.tools[name] = tool
}

func (t *Tools) Get(name string) (Tool, error) {
	t.m.Lock()
	defer t.m.Unlock()

	tool, ok := t.tools[name]
	if ok {
		return tool, nil
	} else {
		return nil, fmt.Errorf("no such tool %s", name)
	}
}

var Default = NewTools()

func Register(name string, tool Tool) {
	Default.Register(name, tool)
}

func Get(name string) (Tool, error) {
	return Default.Get(name)
}
