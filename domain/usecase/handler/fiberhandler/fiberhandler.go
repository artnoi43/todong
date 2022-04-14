package fiberhandler

import (
	"github.com/artnoi43/todong/domain/usecase"
	"github.com/artnoi43/todong/domain/usecase/middleware"
)

type FiberHandler struct {
	dataGateway usecase.DataGateway
	config      *middleware.Config
}

func New(d usecase.DataGateway, c *middleware.Config) *FiberHandler {
	return &FiberHandler{
		dataGateway: d,
		config:      c,
	}
}
