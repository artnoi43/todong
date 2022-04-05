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

	// Use previous values if eq null string
	compareVal := func(old, new string, target *string) {
		var nullString string
		if new == nullString {
			*target = old
			return
		}
		*target = new
	}
	compareStatus := func(old, new enums.Status, target *enums.Status) {
		if new.IsValid() {
			*target = new
		}
		*target = old
	}
	var u datamodel.Todo // Updated to-do
	compareVal(targetTodo.UUID, uuid, &u.UUID)
	compareVal(targetTodo.UserUUID, "", &u.UserUUID)
	compareVal(targetTodo.Title, updatesReq.Title, &u.Title)
	compareVal(targetTodo.Description, updatesReq.Description, &u.Description)
	compareVal(targetTodo.TodoDate, updatesReq.TodoDate, &u.TodoDate)
	compareVal(targetTodo.Image, imgStrReq, &u.Image)
	compareStatus(targetTodo.Status, enums.Status(updatesReq.Status), &u.Status)

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
