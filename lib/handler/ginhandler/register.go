package ginhandler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/artnoi43/todong/datamodel"
	"github.com/artnoi43/todong/enums"
	"github.com/artnoi43/todong/internal"
	"github.com/artnoi43/todong/lib/utils"
)

// Register registers the user in database with bcrypt-hased password.
func (h *GinHandler) Register(c *gin.Context) {
	var req internal.AuthJson
	if err := c.BindJSON(&req); err != nil {
		status := utils.ErrStatus(enums.MapErrHandler.Unmarshal, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, status)
		return
	}
	// Check for existing username in Postgres
	ctx := c.Request.Context()
	var user datamodel.User
	_ = h.dataGateway.GetUserByUsername(ctx, req.Username, &user)
	// If retrived user.Username is not zero-valued,
	// it means the username was used
	if user.Username != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"duplicate username": req.Username,
		})
		return
	}
	pw, err := utils.EncodeBcrypt([]byte(req.Password))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"failed to generate password": err.Error(),
		})
		return
	}
	user = datamodel.User{
		UUID:     uuid.NewString(),
		Username: req.Username,
		Password: pw,
	}
	if err := h.dataGateway.CreateUser(ctx, &user); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("failed to create user: %s", req.Username),
		})
		return
	}
	// user.Password will not leak because the field has "-" in its JSON tag
	c.JSON(http.StatusCreated, gin.H{
		"status":   "user registration successful",
		"username": user.Username,
		"userUuid": user.UUID,
	})
}
