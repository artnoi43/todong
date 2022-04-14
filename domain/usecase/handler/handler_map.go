package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"

	"github.com/artnoi43/todong/domain/usecase/handler/fiberhandler"
	"github.com/artnoi43/todong/domain/usecase/handler/ginhandler"
	"github.com/artnoi43/todong/domain/usecase/handler/httphandler"
	"github.com/artnoi43/todong/lib/enums"
)

// MapHandlers map ginHandler/fiberHandler methods to some strings from enums.
func MapHandlers(
	g *ginhandler.GinHandler,
	f *fiberhandler.FiberHandler,
	m *httphandler.HttpHandler,
) (
	map[enums.Endpoint]func(*gin.Context), // Map for Gin
	map[enums.Endpoint]func(*fiber.Ctx) error, // Map for Fiber
	map[enums.Endpoint]http.HandlerFunc, // Map for Gorilla
) {
	MapGinHandlers := map[enums.Endpoint]func(*gin.Context){
		enums.HandlerRegister:    g.Register,
		enums.HandlerLogin:       g.Login,
		enums.HandlerCreateTodo:  g.CreateTodo,
		enums.HandlerGetTodo:     g.GetTodo,
		enums.HandlerUpdateTodo:  g.UpdateTodo,
		enums.HandlerDeleteTodo:  g.DeleteTodo,
		enums.HandlerNewPassword: g.NewPassword,
		enums.HandlerDeleteUser:  g.DeleteUser,
		enums.HandlerTestAuth:    g.TestAuth,
	}
	MapFiberHandlers := map[enums.Endpoint]func(*fiber.Ctx) error{
		enums.HandlerRegister:    f.Register,
		enums.HandlerLogin:       f.Login,
		enums.HandlerCreateTodo:  f.CreateTodo,
		enums.HandlerGetTodo:     f.GetTodo,
		enums.HandlerUpdateTodo:  f.UpdateTodo,
		enums.HandlerDeleteTodo:  f.DeleteTodo,
		enums.HandlerNewPassword: f.NewPassword,
		enums.HandlerDeleteUser:  f.DeleteUser,
	}
	MapGorillaHandlers := map[enums.Endpoint]http.HandlerFunc{
		enums.HandlerRegister:    m.Register,
		enums.HandlerLogin:       m.Login,
		enums.HandlerCreateTodo:  m.CreateTodo,
		enums.HandlerGetTodo:     m.GetTodo,
		enums.HandlerUpdateTodo:  m.UpdateTodo,
		enums.HandlerDeleteTodo:  m.DeleteTodo,
		enums.HandlerNewPassword: m.NewPassword,
		enums.HandlerDeleteUser:  m.DeleteUser,
	}
	return MapGinHandlers, MapFiberHandlers, MapGorillaHandlers
}
