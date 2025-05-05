package server

import (
	"context"
	"mekanicum/pkg/protocol"
	"sync"
)

type (
	resourcesCapability struct {
		subscribed  bool
		listChanged bool
	}

	promptsCapability struct {
		listChanged bool
	}

	toolsCapability struct {
		listChanged bool
	}

	capabilities struct {
		prompt    *promptsCapability
		resources *resourcesCapability
		tools     *toolsCapability
	}
)

type (
	MCPHandlerFunc func(msg *protocol.McpMessage)
	ContextFunc    func(ctx context.Context) context.Context
)

type Server struct {
	name      string
	version   string
	instruct  string
	prompts   map[string]*protocol.Prompt
	resources map[string]*protocol.Resource
	tools     map[string]*protocol.Tool

	capabilities capabilities

	mu sync.Mutex
}

func New(name, version string, opts ...Option) *Server {
	s := &Server{
		name:    name,
		version: version,

		prompts:   make(map[string]*protocol.Prompt),
		resources: make(map[string]*protocol.Resource),
		tools:     make(map[string]*protocol.Tool),

		capabilities: capabilities{
			prompt:    &promptsCapability{},
			resources: &resourcesCapability{},
			tools:     &toolsCapability{},
		},
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

// TODO: Add RegisterSession
// TODO: Add UnregisterSession
// TODO: Add RegisterTool/s
// TODO: Add RegisterPrompt
// TODO: Add RegisterResources
