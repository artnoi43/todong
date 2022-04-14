package httphandler

import (
	"github.com/artnoi43/todong/domain/usecase"
	"github.com/artnoi43/todong/domain/usecase/middleware"
)

type HttpHandler struct {
	dataGateway usecase.DataGateway
	config      *middleware.Config
}

func New(dataGateway usecase.DataGateway, conf *middleware.Config) *HttpHandler {
	return &HttpHandler{
		dataGateway: dataGateway,
		config:      conf,
	}
}
