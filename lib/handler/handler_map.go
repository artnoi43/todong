package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"

	"github.com/artnoi43/todong/enums"
	"github.com/artnoi43/todong/lib/handler/fiberhandler"
	"github.com/artnoi43/todong/lib/handler/ginhandler"
)

// MapHandlers map ginHandler/fiberHandler methods to some strings from enums.
func MapHandlers(
	g *ginhandler.GinHandler,
	f *fiberhandler.FiberHandler,
) (
	map[string]func(*gin.Context),
	map[string]func(*fiber.Ctx) error,
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
	return MapGinHandlers, MapFiberHandlers
}
