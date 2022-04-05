package ginhandler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/artnoi43/todong/datamodel"
	"github.com/artnoi43/todong/enums"
	"github.com/artnoi43/todong/internal"
	"github.com/artnoi43/todong/lib/utils"
)

// CreateTodo creates a new datamodel.Todo in the database
// Requests to this endpoints have to use 'Content-Type: multipart/form-data'
// Example curl command:
// curl -v -F 'data="{\"key\" = \"val"\}"' -F 'file=@filename' 'http://localhost:8000/todo/create';
func (h *GinHandler) CreateTodo(c *gin.Context) {
	userInfo, err := utils.ExtractAndDecodeJwt(c)
	if err != nil {
		status := utils.ErrStatus(enums.MapErrHandler.JwtError, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, status)
		return
	}
	formData, err := utils.ExtractTodoMultipartFileAndData(c)
	if err != nil && !errors.Is(err, enums.ErrFile) {
		status := utils.ErrStatus(enums.MapErrHandler.MultipartError, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, status)
		return
	}
	var req internal.TodoReqBody
	if err := json.Unmarshal(formData.JSONData, &req); err != nil {
		status := utils.ErrStatus(enums.MapErrHandler.Unmarshal, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, status)
		return
	}
	var imgBytes []byte
	if formData.FileData != nil {
		imgBytes = formData.FileData
	} else {
		// Default to-do image
		imgBytes, err = os.ReadFile("./assets/default.svg")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status": "failed to set default image for todo",
				"error":  fmt.Sprintf("file read error: %v", err.Error()),
			})
		}
	}
	// Convert to Base64 string
	imgBase64Str := base64.StdEncoding.EncodeToString(imgBytes)
	if l := len(imgBase64Str); l > enums.POSTGRES_MAX_STRLEN {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "failed to set image for to-do",
			"error":  fmt.Sprintf("image file too large: %d > %d", l, enums.POSTGRES_MAX_STRLEN),
		})
		return
	}
	todo, err := datamodel.NewTodo(userInfo.UserUuid, req.Title, req.Description, req.TodoDate, enums.InProgress, imgBase64Str)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status": "todo creation failed",
				"error":  err.Error()},
		)
	}

	// Write to store
	if err := h.dataGateway.CreateTodo(c.Request.Context(), todo); err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"status": "todo creation failed",
				"error":  err.Error(),
			},
		)
	}
	c.JSON(http.StatusCreated, todo)
}
