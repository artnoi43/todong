package gorillaserver

import (
	"net/http"
	"time"
)

func (s *gorillaServer) Serve(addr string) error {
	srv := &http.Server{
		Addr:         addr,
		Handler:      s.router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return srv.ListenAndServe()
}
