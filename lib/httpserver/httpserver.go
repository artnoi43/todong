package httpserver

import (
	"context"

	"github.com/artnoi43/todong/enums"
	"github.com/artnoi43/todong/lib/handler"
	"github.com/artnoi43/todong/lib/httpserver/fiberserver"
	"github.com/artnoi43/todong/lib/httpserver/ginserver"
	"github.com/artnoi43/todong/lib/httpserver/gorillaserver"
	"github.com/artnoi43/todong/lib/middleware"
)

// Server abstracts different web frameworks (e.g. Fiber, Gorilla, and Gin)
type Server interface {
	SetUpRoutes(conf *middleware.Config, handler handler.Adaptor)
	Serve(addr string) error
	Shutdown(context.Context)
}

func New(t enums.ServerType) Server {
	if t.IsValid() {
		switch enums.ToUpper(t) {
		case enums.Gin:
			return ginserver.New()
		case enums.Fiber:
			return fiberserver.New()
		case enums.Gorilla:
			return gorillaserver.New()
		}
	}
	panic("invalid server type")
}
