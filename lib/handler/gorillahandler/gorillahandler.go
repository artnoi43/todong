package gorillahandler

import (
	"github.com/artnoi43/todong/lib/middleware"
	"github.com/artnoi43/todong/lib/store"
)

type GorillaHandler struct {
	dataGateway store.DataGateway
	config      *middleware.Config
}

func New(dataGateway store.DataGateway, conf *middleware.Config) *GorillaHandler {
	return &GorillaHandler{
		dataGateway: dataGateway,
		config:      conf,
	}
}
