package fiberhandler

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/artnoi43/todong/data/model"
	"github.com/artnoi43/todong/data/store"
	"github.com/artnoi43/todong/lib/enums"
	"github.com/artnoi43/todong/lib/utils"
)

// GetTodo returns []model.Todo for the users.
// If to-do UUID is given as URL parameter, it returns all of the user's orders.
func (h *FiberHandler) GetTodo(c *fiber.Ctx) error {
	userInfo, err := utils.ExtractAndDecodeJwtFiber(c)
	if err != nil {
		status := utils.ErrStatus(enums.MapErrHandler.JwtError, err)
		return c.Status(http.StatusBadRequest).JSON(status)
	}

	// Get todo-data from DB
	ctx := c.Context()
	var getAll bool
	var targetTodo model.Todo
	var targetTodos []model.Todo
	// If UUID is not given as URL parameter, find all todo records for user
	uuid := c.Params("uuid")
	if len(uuid) == 0 {
		getAll = true
		err = h.dataGateway.GetUserTodos(ctx, &model.Todo{
			UserUUID: userInfo.UserUuid,
		}, &targetTodos)
	} else {
		err = h.dataGateway.GetOneTodo(ctx, &model.Todo{
			UserUUID: userInfo.UserUuid,
			UUID:     uuid,
		}, &targetTodo)
	}
	if err != nil {
		// Record not found
		if errors.Is(err, store.ErrRecordNotFound) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"error": "todo not found",
				"uuid":  uuid,
			})
		}
		// Other errors
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": errors.New("failed to get todo"),
		})
	}
	if getAll {
		return c.Status(http.StatusOK).JSON(targetTodos)
	}
	return c.Status(http.StatusOK).JSON(targetTodo)
}
