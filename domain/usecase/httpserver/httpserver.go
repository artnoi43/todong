package httpserver

import (
	"context"

	"github.com/artnoi43/mgl/str"

	"github.com/artnoi43/todong/domain/usecase/handler"
	"github.com/artnoi43/todong/domain/usecase/httpserver/fiberserver"
	"github.com/artnoi43/todong/domain/usecase/httpserver/ginserver"
	"github.com/artnoi43/todong/domain/usecase/httpserver/gorillaserver"
	"github.com/artnoi43/todong/domain/usecase/httpserver/nethttpserver"
	"github.com/artnoi43/todong/domain/usecase/middleware"
	"github.com/artnoi43/todong/lib/enums"
)

// Server abstracts different web frameworks (e.g. Fiber, Gorilla, and Gin)
type Server interface {
	SetUpRoutes(*middleware.Config, handler.Adapter)
	Serve(addr string) error
	Shutdown(context.Context)
}

func New(t enums.ServerType) Server {
	if t.IsValid() {
		switch str.ToUpper(t) {
		case enums.Gin:
			return ginserver.New()
		case enums.Fiber:
			return fiberserver.New()
		case enums.Gorilla:
			return gorillaserver.New()
		case enums.NetHttp:
			return nethttpserver.New()
		}
	}
	panic("invalid server type")
}
