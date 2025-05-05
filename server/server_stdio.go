package server

import (
	"context"
	"io"
	"log"
	"mekanicum/pkg/protocol"
	"os"
	"sync/atomic"
)

type stdioSession struct {
	initialized atomic.Bool
	notifyCh    chan protocol.McpMessage
}

func (session *stdioSession) ID() string {
	return "stdio-session"
}

func (session *stdioSession) Initialize() {
	session.initialized.Store(true)
}

func (session *stdioSession) Initialized() bool {
	return session.initialized.Load()
}

func (session *stdioSession) Notifications() chan<- protocol.McpMessage {
	return session.notifyCh
}

type StdioServer struct {
	ctxFn  ContextFunc
	errLog *log.Logger
	server *Server
}

func NewStdinServer(server Server) *StdioServer {
	s := &StdioServer{
		server: &server,
	}

	return s
}

func (s *StdioServer) SetContextFn(fn ContextFunc) {
	s.ctxFn = fn
}

func (s *StdioServer) Serve() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	return s.listen(ctx, os.Stdin, os.Stdout)
}

func (s *StdioServer) listen(ctx context.Context, in io.Reader, out io.Writer) error {
	// TODO: Handle session registration here

	// TODO: Handle context overriding here

	return nil
}
