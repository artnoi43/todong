package fiberhandler

import (
	"errors"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/artnoi43/todong/data/model"
	"github.com/artnoi43/todong/data/store"
	"github.com/artnoi43/todong/internal"
	"github.com/artnoi43/todong/lib/enums"
	"github.com/artnoi43/todong/lib/utils"
)

// Login authenticates username/password and return JWT token signed with configured secret
func (h *FiberHandler) Login(c *fiber.Ctx) error {
	var req internal.AuthJson
	if err := c.BodyParser(&req); err != nil {
		status := utils.ErrStatus(enums.MapErrHandler.Unmarshal, err)
		return c.Status(http.StatusBadRequest).JSON(status)
	}

	ctx := c.Context()
	var user model.User
	if err := h.dataGateway.GetUserByUsername(ctx, req.Username, &user); err != nil {
		if !errors.Is(err, store.ErrRecordNotFound) {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "login failed",
			})
		}
	}
	// Null user from database, i.e. zero-valued user.UUID
	if user.UUID == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status": "invalid username or password",
			"code":   1,
		})
	}
	// Compare hashed password with bcrypt
	if err := utils.DecodeBcrypt(user.Password, []byte(req.Password)); err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"status": "invalid username or password",
			"code":   2,
		})
	}
	token, exp, err := utils.NewJwtToken(user.UUID, []byte(h.config.SecretKey))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "login failed",
		})
	}
	resp := internal.LoginResponse(struct {
		Status   string
		Username string
		UserUuid string
		Exp      time.Time
		Token    string
	}{
		Status:   "login successful",
		Username: user.Username,
		UserUuid: user.UUID,
		Exp:      exp,
		Token:    token,
	})
	return c.Status(http.StatusOK).JSON(resp)
}
