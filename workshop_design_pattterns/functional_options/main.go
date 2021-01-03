package main

import (
	"fmt"
	"log"
)

type ServerConfig struct {
	Hostname string
	Port     string
}

func (s *ServerConfig) Address() string {
	return fmt.Sprintf("%s:%s", s.Hostname, s.Port)
}


type ServerOptions func(*ServerConfig)

func WithHostname(hostname string) ServerOptions {
	return func(s *ServerConfig) {
		s.Hostname = hostname
	}
}

func WithPort(port string) ServerOptions {
	return func(s *ServerConfig) {
		s.Port = port
	}
}

func NewServer(opts ...ServerOptions) *ServerConfig {
	s := &ServerConfig{
		Hostname: "localhost",
		Port:     "80",
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}
func main() {
	log.Printf("Creating server config")
	server := NewServer(
		WithHostname("144.22.22.1"),
		WithPort("443"),
	)

	log.Printf(server.Address())

}

// OUTPUT:
// 2021/01/03 00:24:05 Creating server config
// 2021/01/03 00:24:05 144.22.22.1:443
