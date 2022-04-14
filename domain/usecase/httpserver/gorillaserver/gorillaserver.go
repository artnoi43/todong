package gorillaserver

import (
	"context"

	"github.com/gorilla/mux"
)

type gorillaServer struct {
	router *mux.Router
}

func New() *gorillaServer {
	return &gorillaServer{
		router: mux.NewRouter(),
	}
}

// Implements httpserver.Server
func (s *gorillaServer) Shutdown(ctx context.Context) {

}
