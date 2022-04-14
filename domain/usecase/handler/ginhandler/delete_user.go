package ginhandler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/artnoi43/todong/data/model"
	"github.com/artnoi43/todong/data/store"
	"github.com/artnoi43/todong/lib/enums"
	"github.com/artnoi43/todong/lib/utils"
)

// DeleteUser deletes a model.User in database
// model.User.UUID is used to target deletion
func (h *GinHandler) DeleteUser(c *gin.Context) {
	userInfo, err := utils.ExtractAndDecodeJwt(c)
	if err != nil {
		status := enums.MapErrHandler.JwtError
		status["error"] = err.Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, status)
		return
	}
	uuid := userInfo.UserUuid
	ctx := c.Request.Context()
	// Delete data from DB
	if err = h.dataGateway.DeleteUser(ctx, &model.User{
		UUID: uuid,
	}); err != nil {
		if errors.Is(err, store.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status": "user not found",
			})
			return
		}
		// Other errors
		c.AbortWithStatusJSON(http.StatusInternalServerError, errors.New("failed to delete user"))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "user deletion successful",
		"uuid":   uuid,
	})
}
