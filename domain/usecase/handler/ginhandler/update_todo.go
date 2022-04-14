package ginhandler

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/artnoi43/todong/data/model"
	"github.com/artnoi43/todong/data/store"
	"github.com/artnoi43/todong/internal"
	"github.com/artnoi43/todong/lib/enums"
	"github.com/artnoi43/todong/lib/utils"
)

// UpdateTodo updates user's model.Todo in database
func (h *GinHandler) UpdateTodo(c *gin.Context) {
	uuid := c.Param("uuid")
	userInfo, err := utils.ExtractAndDecodeJwt(c)
	if err != nil {
		status := utils.ErrStatus(enums.MapErrHandler.JwtError, err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, status)
		return
	}

	// Find targetTodo in database
	ctx := c.Request.Context()
	var targetTodo model.Todo
	if err := h.dataGateway.GetOneTodo(ctx, &model.Todo{
		UserUUID: userInfo.UserUuid,
		UUID:     uuid,
	}, &targetTodo); err != nil {
		if errors.Is(err, store.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "todo not found",
				"uuid":  uuid,
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Extract multipart values from keys "file" and "data"
	formData, err := utils.ExtractTodoMultipartFileAndData(c)
	// utils.ErrFile is soft error
	if err != nil && !errors.Is(err, enums.ErrFile) {
		status := utils.ErrStatus(enums.MapErrHandler.MultipartError, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, status)
		return
	}
	// Continue if soft errors
	var updatesReq internal.TodoReqBody
	if err := json.Unmarshal(formData.JSONData, &updatesReq); err != nil {
		status := utils.ErrStatus(enums.MapErrHandler.Unmarshal, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, status)
		return
	}
	var imgStrReq string
	// If image file was uploaded, encode it to Base64
	if formData.FileData != nil {
		imgStrReq = base64.StdEncoding.EncodeToString(formData.FileData)
	}
	if l := len(imgStrReq); l > enums.POSTGRES_MAX_STRLEN {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   fmt.Sprintf("image file too large: %d", l),
			"maximum": enums.POSTGRES_MAX_STRLEN,
		})
		return
	}

	var u model.Todo // This to-do will be written to data store
	utils.UpdatedTodo(uuid, imgStrReq, &updatesReq, &targetTodo, &u)

	// Update data in DB
	if err := h.dataGateway.UpdateTodo(ctx, &model.Todo{
		UserUUID: userInfo.UserUuid,
		UUID:     uuid,
	}, &u); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update todo",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":   "todo update successful",
		"uuid":     u.UUID,
		"userUuid": u.UserUUID,
	})
}
