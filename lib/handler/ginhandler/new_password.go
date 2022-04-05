package ginhandler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/artnoi43/todong/datamodel"
	"github.com/artnoi43/todong/enums"
	"github.com/artnoi43/todong/internal"
	"github.com/artnoi43/todong/lib/store"
	"github.com/artnoi43/todong/lib/utils"
)

// DeleteUser deletes a datamodel.User in database
// datamodel.User.UUID is used to target deletion
func (h *GinHandler) NewPassword(c *gin.Context) {
	userInfo, err := utils.ExtractAndDecodeJwt(c)
	if err != nil {
		status := utils.ErrStatus(enums.MapErrHandler.JwtError, err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, status)
		return
	}
	uuid := userInfo.UserUuid

	// Parse new password JSON request
	var newPassReq internal.NewPasswordJson
	if err := c.BindJSON(&newPassReq); err != nil {
		status := utils.ErrStatus(enums.MapErrHandler.Unmarshal, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, status)
		return
	}
	if len(newPassReq.NewPassword) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "blank password received",
		})
		return
	}

	// Get user from DB
	ctx := c.Request.Context()
	var targetUser datamodel.User
	if err := h.dataGateway.GetUserByUuid(ctx, uuid, &targetUser); err != nil {
		// Record not found
		if errors.Is(err, store.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status": "user not found",
				"uuid":   uuid,
			})
			return
		}
		// Other errors
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "failed to find target user",
			"uuid":   uuid,
			"error":  err.Error(),
		})
		return
	}
	pw, err := utils.EncodeBcrypt([]byte(newPassReq.NewPassword))
	if err != nil {
		if errors.Is(enums.ErrPwTooShort, err) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "failed to change password",
		})
		return
	}
	// Update data in DB
	if err := h.dataGateway.ChangePassword(ctx, &targetUser, pw); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"errors": "failed to change password",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   "password change successful",
		"username": targetUser.Username,
	})
}
