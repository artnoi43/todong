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

// GetTodo returns []datamodel.Todo for the users.
// If to-do UUID is given as URL parameter, it returns all of the user's orders.
func (h *GinHandler) GetTodo(c *gin.Context) {
	uuid := c.Param("uuid")
	userInfo, err := utils.ExtractAndDecodeJwt(c)
	if err != nil {
		status := utils.ErrStatus(enums.MapErrHandler.JwtError, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, status)
		return
	}

	// Get todo-data from DB
	ctx := c.Request.Context()
	var getAll bool
	var targetTodo datamodel.Todo
	var targetTodos []datamodel.Todo
	// If UUID is not given as URL parameter, find all todo records for user
	if len(uuid) == 0 {
		getAll = true
		err = h.dataGateway.GetUserTodos(ctx, &datamodel.Todo{
			UserUUID: userInfo.UserUuid,
		}, &targetTodos)
	} else {
		err = h.dataGateway.GetOneTodo(ctx, &datamodel.Todo{
			UserUUID: userInfo.UserUuid,
			UUID:     uuid,
		}, &targetTodo)
	}
	if err != nil {
		// Record not found
		if errors.Is(err, store.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "todo not found",
				"uuid":  uuid,
			})
			return
		}
		// Other errors
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": errors.New("failed to get todo"),
		})
		return
	}
	if getAll {
		c.JSON(http.StatusOK, targetTodos)
		return
	}
	c.JSON(http.StatusOK, targetTodo)
}
