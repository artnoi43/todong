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

// DeleteTodo deletes a model.Todo in database
// To-do UUID is used to target deletion
func (h *FiberHandler) DeleteTodo(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	userInfo, err := utils.ExtractAndDecodeJwtFiber(c)
	if err != nil {
		status := utils.ErrStatus(enums.MapErrHandler.JwtError, err)
		return c.Status(http.StatusBadRequest).JSON(status)
	}

	// Delete data from DB
	ctx := c.Context()
	var targetTodo model.Todo
	err = h.dataGateway.GetOneTodo(ctx, &model.Todo{
		UserUUID: userInfo.UserUuid,
		UUID:     uuid,
	}, &targetTodo)
	if err != nil {
		// Record not found
		if errors.Is(err, store.ErrRecordNotFound) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status": "todo not found",
				"uuid":   uuid,
			})
		}
		// Other errors
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": "failed to delete todo",
			"error":  err.Error(),
		})

	}
	if err = h.dataGateway.DeleteTodo(ctx, &model.Todo{
		UserUUID: userInfo.UserUuid,
		UUID:     uuid,
	}); err != nil {
		// Other errors
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": "failed to delete todo",
			"uuid":   uuid,
			"error":  err.Error(),
		})

	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "todo deletion successful",
		"uuid":   uuid,
	})
}
