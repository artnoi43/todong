package ginhandler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/artnoi43/todong/datamodel"
	"github.com/artnoi43/todong/enums"
	"github.com/artnoi43/todong/lib/store"
	"github.com/artnoi43/todong/lib/utils"
)

// DeleteTodo deletes a datamodel.Todo in database
// To-do UUID is used to target deletion
func (h *GinHandler) DeleteTodo(c *gin.Context) {
	uuid := c.Param("uuid")
	userInfo, err := utils.ExtractAndDecodeJwt(c)
	if err != nil {
		status := utils.ErrStatus(enums.MapErrHandler.JwtError, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, status)
		return
	}

	// Delete data from DB
	ctx := c.Request.Context()
	var targetTodo datamodel.Todo
	err = h.dataGateway.GetOneTodo(ctx, &datamodel.Todo{
		UserUUID: userInfo.UserUuid,
		UUID:     uuid,
	}, &targetTodo)
	if err != nil {
		// Record not found
		if errors.Is(err, store.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status": "todo not found",
				"uuid":   uuid,
			})
			return
		}
		// Other errors
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "failed to delete todo",
			"error":  err.Error(),
		})
		return
	}
	if err = h.dataGateway.DeleteTodo(ctx, &datamodel.Todo{
		UserUUID: userInfo.UserUuid,
		UUID:     uuid,
	}); err != nil {
		// Other errors
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "failed to delete todo",
			"uuid":   uuid,
			"error":  err.Error(),
		})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"status": "todo deletion successful",
		"uuid":   uuid,
	})
}
