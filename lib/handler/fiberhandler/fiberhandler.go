package fiberhandler

import (
	"github.com/artnoi43/todong/lib/middleware"
	"github.com/artnoi43/todong/lib/store"
)

type FiberHandler struct {
	dataGateway store.DataGateway
	config      *middleware.Config
}

func New(d store.DataGateway, c *middleware.Config) *FiberHandler {
	return &FiberHandler{
		dataGateway: d,
		config:      c,
	}
}
