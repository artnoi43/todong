package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"

	"github.com/artnoi43/todong/enums"
	"github.com/artnoi43/todong/lib/handler/fiberhandler"
	"github.com/artnoi43/todong/lib/handler/ginhandler"
	"github.com/artnoi43/todong/lib/handler/gorillahandler"
	"github.com/artnoi43/todong/lib/middleware"
	"github.com/artnoi43/todong/lib/store"
)

// Adapter abstracts handlers of different web frameworks
type Adaptor interface {
	// These methods takes path as param,
	// and returns a matching handler func from adapter maps.
	Gin(enums.Endpoint) func(*gin.Context)
	Fiber(enums.Endpoint) func(*fiber.Ctx) error
	Gorilla(enums.Endpoint) http.HandlerFunc
}

// adapter implements Adapter
type adapter struct {
	gin        *ginhandler.GinHandler
	fiber      *fiberhandler.FiberHandler
	gorilla    *gorillahandler.GorillaHandler
	ginMap     map[enums.Endpoint]func(*gin.Context)
	fiberMap   map[enums.Endpoint]func(*fiber.Ctx) error
	gorillaMap map[enums.Endpoint]http.HandlerFunc
}

func (h *adapter) Gin(s enums.Endpoint) func(*gin.Context) {
	if h.ginMap[s] == nil {
		panic(fmt.Sprintf("missing gin handlers for: %s", s))
	}
	return h.ginMap[s]
}

func (h *adapter) Fiber(s enums.Endpoint) func(*fiber.Ctx) error {
	if h.fiberMap[s] == nil {
		panic(fmt.Sprintf("missing fiber handlers for: %s", s))
	}
	return h.fiberMap[s]
}

func (h *adapter) Gorilla(s enums.Endpoint) http.HandlerFunc {
	if h.gorillaMap[s] == nil {
		panic(fmt.Sprintf("missing gorilla/mux handlers for: %s", s))
	}
	return h.gorillaMap[s]
}

func NewAdaptor(t enums.ServerType, conf *middleware.Config, dataGateway store.DataGateway) Adaptor {
	var g *ginhandler.GinHandler
	var f *fiberhandler.FiberHandler
	var m *gorillahandler.GorillaHandler
	switch enums.ToUpper(t) {
	case enums.Gin:
		g = ginhandler.New(dataGateway, conf)
		mapGin, _, _ := MapHandlers(g, f, m)
		return &adapter{
			gin:    g,
			ginMap: mapGin,
		}
	case enums.Fiber:
		f = fiberhandler.New(dataGateway, conf)
		_, mapFiber, _ := MapHandlers(g, f, m)
		return &adapter{
			fiber:    f,
			fiberMap: mapFiber,
		}
	case enums.Gorilla:
		m = gorillahandler.New(dataGateway, conf)
		_, _, mapGorilla := MapHandlers(g, f, m)
		return &adapter{
			gorilla:    m,
			gorillaMap: mapGorilla,
		}
	}
	panic("invalid server type")
}
