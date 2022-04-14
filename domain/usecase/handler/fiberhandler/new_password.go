package fiberhandler

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/artnoi43/todong/data/model"
	"github.com/artnoi43/todong/data/store"
	"github.com/artnoi43/todong/internal"
	"github.com/artnoi43/todong/lib/enums"
	"github.com/artnoi43/todong/lib/utils"
)

// DeleteUser deletes a model.User in database
// model.User.UUID is used to target deletion
func (h *FiberHandler) NewPassword(c *fiber.Ctx) error {
	userInfo, err := utils.ExtractAndDecodeJwtFiber(c)
	if err != nil {
		status := utils.ErrStatus(enums.MapErrHandler.JwtError, err)
		return c.Status(http.StatusInternalServerError).JSON(status)
	}
	uuid := userInfo.UserUuid

	// Parse new password JSON request
	var newPassReq internal.NewPasswordJson
	if err := c.BodyParser(&newPassReq); err != nil {
		status := utils.ErrStatus(enums.MapErrHandler.Unmarshal, err)
		return c.Status(http.StatusBadRequest).JSON(status)
	}
	if len(newPassReq.NewPassword) == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "blank password received",
		})
	}

	// Get user from DB
	ctx := c.Context()
	var targetUser model.User
	if err := h.dataGateway.GetUserByUuid(ctx, uuid, &targetUser); err != nil {
		// Record not found
		if errors.Is(err, store.ErrRecordNotFound) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status": "user not found",
				"uuid":   uuid,
			})
		}
		// Other errors
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": "failed to find target user",
			"uuid":   uuid,
			"error":  err.Error(),
		})
	}
	pw, err := utils.EncodeBcrypt([]byte(newPassReq.NewPassword))
	if err != nil {
		if errors.Is(enums.ErrPwTooShort, err) {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to change password",
		})
	}
	// Update data in DB
	if err := h.dataGateway.ChangePassword(ctx, &targetUser, pw); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"errors": "failed to change password",
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":   "password change successful",
		"username": targetUser.Username,
	})
}
