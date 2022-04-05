package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"

	"github.com/artnoi43/todong/enums"
	"github.com/artnoi43/todong/lib/handler/fiberhandler"
	"github.com/artnoi43/todong/lib/handler/ginhandler"
	"github.com/artnoi43/todong/lib/handler/gorillahandler"
)

// MapHandlers map ginHandler/fiberHandler methods to some strings from enums.
func MapHandlers(
	g *ginhandler.GinHandler,
	f *fiberhandler.FiberHandler,
	m *gorillahandler.GorillaHandler,
) (
	map[string]func(*gin.Context),
	map[string]func(*fiber.Ctx) error,
	map[string]http.HandlerFunc,
) {
	MapGinHandlers := map[string]func(*gin.Context){
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
	MapFiberHandlers := map[string]func(*fiber.Ctx) error{
		enums.HandlerRegister:    f.Register,
		enums.HandlerLogin:       f.Login,
		enums.HandlerCreateTodo:  f.CreateTodo,
		enums.HandlerGetTodo:     f.GetTodo,
		enums.HandlerUpdateTodo:  f.UpdateTodo,
		enums.HandlerDeleteTodo:  f.DeleteTodo,
		enums.HandlerNewPassword: f.NewPassword,
		enums.HandlerDeleteUser:  f.DeleteUser,
	}
	MapGorillaHandlers := map[string]http.HandlerFunc{
		enums.HandlerRegister:   m.Register,
		enums.HandlerLogin:      m.Login,
		enums.HandlerCreateTodo: m.CreateTodo,
		enums.HandlerGetTodo:    m.GetTodo,
	}
	return MapGinHandlers, MapFiberHandlers, MapGorillaHandlers
}
