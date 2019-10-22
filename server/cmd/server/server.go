package main

import (
	"context"
	"fmt"
	"github.com/Tornado9966/menu-example/server/menu"
	"net/http"
)

type HttpPortNumber int

// MenuApiServer configures necessary handlers and starts listening on a configured port.
type MenuApiServer struct {
	Port HttpPortNumber

	DishesHandler menu.HttpHandlerFunc

	server *http.Server
}

// Start will set all handlers and start listening.
// If this methods succeeds, it does not return until server is shut down.
// Returned error will never be nil.
func (s *MenuApiServer) Start() error {
	if s.DishesHandler == nil {
		return fmt.Errorf("dishes HTTP handler is not defined - cannot start")
	}
	if s.Port == 0 {
		return fmt.Errorf("port is not defined")
	}

	handler := new(http.ServeMux)
	handler.HandleFunc("/dishes", s.DishesHandler)

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: handler,
	}

	return s.server.ListenAndServe()
}

// Stops will shut down previously started HTTP server.
func (s *MenuApiServer) Stop() error {
	if s.server == nil {
		return fmt.Errorf("server was not started")
	}
	return s.server.Shutdown(context.Background())
}
