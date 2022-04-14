package ginhandler

import (
	"github.com/artnoi43/todong/domain/usecase"
	"github.com/artnoi43/todong/domain/usecase/middleware"
)

type GinHandler struct {
	dataGateway usecase.DataGateway
	config      *middleware.Config
}

func New(d usecase.DataGateway, c *middleware.Config) *GinHandler {
	return &GinHandler{
		dataGateway: d,
		config:      c,
	}
}
