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

func (h *FiberHandler) DeleteUser(c *fiber.Ctx) error {
	userInfo, err := utils.ExtractAndDecodeJwtFiber(c)
	if err != nil {
		status := enums.MapErrHandler.JwtError
		status["error"] = err.Error()
		return c.Status(http.StatusInternalServerError).JSON(status)
	}
	uuid := userInfo.UserUuid
	ctx := c.Context()
	// Delete data from DB
	if err = h.dataGateway.DeleteUser(ctx, &model.User{
		UUID: uuid,
	}); err != nil {
		if errors.Is(err, store.ErrRecordNotFound) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status": "user not found",
			})
		}
		// Other errors
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to delete user",
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "user deletion successful",
		"uuid":   uuid,
	})
}
