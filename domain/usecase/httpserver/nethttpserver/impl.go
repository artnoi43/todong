package nethttpserver

import (
	"context"
	"net/http"

	"github.com/artnoi43/todong/domain/usecase/handler"
	"github.com/artnoi43/todong/domain/usecase/middleware"
	"github.com/artnoi43/todong/lib/utils"
)

func (s *nethttpServer) Shutdown(ctx context.Context) {
	// TODO: implement
}

// // Routes handling is done in *nethttpServer.ServeHTTP
// // i.e., after *nethttpServer.Serve is called
func (s *nethttpServer) SetUpRoutes(conf *middleware.Config, adapter handler.Adapter) {
	s.auth = utils.NewAuthenticator([]byte(conf.SecretKey))
	s.adapter = adapter
}

func (s *nethttpServer) Serve(addr string) error {
	return http.ListenAndServe(addr, s)
}
