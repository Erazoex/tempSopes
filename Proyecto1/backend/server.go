package main

import (
	"backend/routes"
	"fmt"
	"net/http"
)

type Server struct {
	HttpServer http.Server
}

func NewServer(address string) *Server {
	routes.Init()
	return &Server{
		HttpServer: http.Server{
			Addr: address,
		},
	}
}

func (s *Server) ListenAndServe() error {
	fmt.Printf("Listenning to port %s\n", s.HttpServer.Addr)
	return s.HttpServer.ListenAndServe()
}
