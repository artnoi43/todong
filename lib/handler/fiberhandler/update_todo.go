package fiberhandler

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/artnoi43/todong/datamodel"
	"github.com/artnoi43/todong/enums"
	"github.com/artnoi43/todong/internal"
	"github.com/artnoi43/todong/lib/store"
	"github.com/artnoi43/todong/lib/utils"
)

// UpdateTodo updates user's datamodel.Todo in database
func (h *FiberHandler) UpdateTodo(c *fiber.Ctx) error {
	userInfo, err := utils.ExtractAndDecodeJwtFiber(c)
	if err != nil {
		status := utils.ErrStatus(enums.MapErrHandler.JwtError, err)
		return c.Status(http.StatusInternalServerError).JSON(status)
	}
	uuid := c.Params("uuid")

	// Find targetTodo in database
	ctx := c.Context()
	var targetTodo datamodel.Todo
	if err := h.dataGateway.GetOneTodo(ctx, &datamodel.Todo{
		UserUUID: userInfo.UserUuid,
		UUID:     uuid,
	}, &targetTodo); err != nil {
		if errors.Is(err, store.ErrRecordNotFound) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"error": "todo not found",
				"uuid":  uuid,
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Extract multipart values from keys "file" and "data"
	formData, err := utils.ExtractTodoMultipartFileAndDataFiber(c)
	// utils.ErrFile is soft error
	if err != nil && !errors.Is(err, enums.ErrFile) {
		status := utils.ErrStatus(enums.MapErrHandler.MultipartError, err)
		return c.Status(http.StatusBadRequest).JSON(status)
	}
	// Continue if soft errors
	var updatesReq internal.TodoReqBody
	if err := json.Unmarshal(formData.JSONData, &updatesReq); err != nil {
		status := utils.ErrStatus(enums.MapErrHandler.Unmarshal, err)
		return c.Status(http.StatusBadRequest).JSON(status)
	}
	var imgStrReq string
	// If image file was uploaded, encode it to Base64
	if formData.FileData != nil {
		imgStrReq = base64.StdEncoding.EncodeToString(formData.FileData)
	}
	if l := len(imgStrReq); l > enums.POSTGRES_MAX_STRLEN {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":   fmt.Sprintf("image file too large: %d", l),
			"maximum": enums.POSTGRES_MAX_STRLEN,
		})
	}

	// TODO: maybe abstract this whole mess
	var u datamodel.Todo // Updated to-do
	utils.UpdatedTodo(uuid, imgStrReq, &updatesReq, &targetTodo, &u)

	// Update data in DB
	if err := h.dataGateway.UpdateTodo(ctx, &datamodel.Todo{
		UserUUID: userInfo.UserUuid,
		UUID:     uuid,
	}, &u); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to update todo",
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status":   "todo update successful",
		"uuid":     u.UUID,
		"userUuid": u.UserUUID,
	})
}
