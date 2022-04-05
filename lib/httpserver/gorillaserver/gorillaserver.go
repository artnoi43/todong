package gorillaserver

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type gorillaServer struct {
	router *mux.Router
	server *http.Server
}

func New() *gorillaServer {
	// r is pointer to the same mux.Router
	// in both gorillaServer.router
	// and gorillaServer.server.router
	r := mux.NewRouter()
	return &gorillaServer{
		router: r,
		server: &http.Server{
			Handler:      r,
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		},
	}
}

func (s *gorillaServer) Shutdown(ctx context.Context) {

}
