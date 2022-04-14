package ginhandler

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/artnoi43/todong/data/model"
	"github.com/artnoi43/todong/data/store"
	"github.com/artnoi43/todong/internal"
	"github.com/artnoi43/todong/lib/enums"
	"github.com/artnoi43/todong/lib/utils"
)

// Login authenticates username/password and return JWT token signed with configured secret
func (h *GinHandler) Login(c *gin.Context) {
	var req internal.AuthJson
	if err := c.BindJSON(&req); err != nil {
		status := utils.ErrStatus(enums.MapErrHandler.Unmarshal, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, status)
		return
	}

	ctx := c.Request.Context()
	var user model.User
	if err := h.dataGateway.GetUserByUsername(ctx, req.Username, &user); err != nil {
		if !errors.Is(err, store.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "login failed",
			})
			return
		}
	}
	// Null user from database, i.e. zero-valued user.UUID
	if user.UUID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "invalid username or password",
			"code":   1,
		})
		return
	}
	// Compare hashed password with bcrypt
	if err := utils.DecodeBcrypt(user.Password, []byte(req.Password)); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": "invalid username or password",
			"code":   2,
		})
		return
	}
	token, exp, err := utils.NewJwtToken(user.UUID, []byte(h.config.SecretKey))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "failed to generate token",
			"error":  err.Error(),
		})
		return
	}
	resp := internal.LoginResponse(struct {
		Status   string
		Username string
		UserUuid string
		Exp      time.Time
		Token    string
	}{
		Status:   "login successful",
		Username: user.Username,
		UserUuid: user.UUID,
		Exp:      exp,
		Token:    token,
	})
	c.JSON(http.StatusOK, resp)
}
