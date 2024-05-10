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

type Server struct {
	host    string
	port    int
	timeout time.Duration
	maxConn int
}

func New(options ...func(*Server)) *Server {
	svr := &Server{}
	for _, opt := range options {
		opt(svr)
	}
	return svr
}

func (s *Server) Start() error {
	// todo
	return nil
}

func WithHost(host string) func(*Server) {
	return func(s *Server) {
		s.host = host
	}
}

func WithPort(port int) func(*Server) {
	return func(s *Server) {
		s.port = port
	}
}

func WithTimeout(timeout time.Duration) func(*Server) {
	return func(s *Server) {
		s.timeout = timeout
	}
}

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
