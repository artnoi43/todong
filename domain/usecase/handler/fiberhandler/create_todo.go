package fiberhandler

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"

	"github.com/artnoi43/todong/data/model"
	"github.com/artnoi43/todong/internal"
	"github.com/artnoi43/todong/lib/enums"
	"github.com/artnoi43/todong/lib/utils"
)

// CreateTodo creates a new model.Todo in the database
// Requests to this endpoints have to use 'Content-Type: multipart/form-data'
// Example curl command:
// curl -v -F 'data="{\"key\" = \"val"\}"' -F 'file=@filename' 'http://localhost:8000/todo/create';
func (h *FiberHandler) CreateTodo(c *fiber.Ctx) error {
	userInfo, err := utils.ExtractAndDecodeJwtFiber(c)
	if err != nil {
		status := utils.ErrStatus(enums.MapErrHandler.JwtError, err)
		return c.Status(http.StatusBadRequest).JSON(status)
	}
	formData, err := utils.ExtractTodoMultipartFileAndDataFiber(c)
	if err != nil && !errors.Is(err, enums.ErrFile) {
		status := utils.ErrStatus(enums.MapErrHandler.MultipartError, err)
		return c.Status(http.StatusBadRequest).JSON(status)
	}
	var req internal.TodoReqBody
	if err := json.Unmarshal(formData.JSONData, &req); err != nil {
		status := utils.ErrStatus(enums.MapErrHandler.Unmarshal, err)
		return c.Status(http.StatusBadRequest).JSON(status)
	}
	var imgBytes []byte
	if formData.FileData != nil {
		imgBytes = formData.FileData
	} else {
		// Default to-do image
		imgBytes, err = os.ReadFile("./assets/default.svg")
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"status": "failed to set default image for todo",
				"error":  fmt.Sprintf("file read error: %v", err.Error()),
			})
		}
	}
	// Convert to Base64 string
	imgBase64Str := base64.StdEncoding.EncodeToString(imgBytes)
	if l := len(imgBase64Str); l > enums.POSTGRES_MAX_STRLEN {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": "failed to set image for to-do",
			"error":  fmt.Sprintf("image file too large: %d > %d", l, enums.POSTGRES_MAX_STRLEN),
		})
	}
	todo, err := model.NewTodo(userInfo.UserUuid, req.Title, req.Description, req.TodoDate, enums.InProgress, imgBase64Str)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status": "todo creation failed",
			"error":  err.Error()},
		)
	}

	// Write to store
	if err := h.dataGateway.CreateTodo(c.Context(), todo); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": "todo creation failed",
			"error":  err.Error(),
		})
	}
	return c.Status(http.StatusCreated).JSON(todo)
}
