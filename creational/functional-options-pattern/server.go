package main

import (
	"log"
	"time"
)

/*
// From a custom Config struct
type Server struct {
	cfg Config
}

type Config struct {
	host    string
	port    int
	timeout time.Duration
	maxConn int
}

func New(cfg Config) *Server {
	return &Server{cfg}
}

func (s *Server) Start() error {
	// todo
	return nil
}

func main() {
	svr := New(Config{"localhost", 1234, 30 * time.Second, 10})
	if err := svr.Start(); err != nil {
		log.Fatal(err)
	}
}
*/

// To Functional Options Pattern

// Server contains the server configuration.
type Server struct {
	host    string
	port    int
	timeout time.Duration
	maxConn int
}

// New creates a server instance with custom options.
func New(options ...func(*Server)) *Server {
	svr := &Server{}
	for _, opt := range options {
		opt(svr)
	}
	return svr
}

// Start runs the server
func (s *Server) Start() error {
	// todo
	return nil
}

// WithHost configures the address fot the server.
func WithHost(host string) func(*Server) {
	return func(s *Server) {
		s.host = host
	}
}

// WithPort configures the port for the server.
func WithPort(port int) func(*Server) {
	return func(s *Server) {
		s.port = port
	}
}

// WithTimeout configures the timeout for the server.
func WithTimeout(timeout time.Duration) func(*Server) {
	return func(s *Server) {
		s.timeout = timeout
	}
}

// WithMacConn configures the maximum number of connections to the server.
func WithMaxConn(maxConn int) func(*Server) {
	return func(s *Server) {
		s.maxConn = maxConn
	}
}

func main() {
	svr := New(
		WithHost("localhost"),
		WithPort(8080),
		WithTimeout(time.Minute),
		WithMaxConn(120),
	)
	if err := svr.Start(); err != nil {
		log.Fatal(err)
	}
}
