package ginhandler

import (
	"github.com/artnoi43/todong/lib/middleware"
	"github.com/artnoi43/todong/lib/store"
)

type GinHandler struct {
	dataGateway store.DataGateway
	config      *middleware.Config
}

func New(d store.DataGateway, c *middleware.Config) *GinHandler {
	return &GinHandler{
		dataGateway: d,
		config:      c,
	}
}
