package ginhandler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/artnoi43/todong/lib/utils"
)

func (h *GinHandler) TestAuth(c *gin.Context) {
	userInfo, err := utils.ExtractAndDecodeJwt(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error: failed to extract jwt: %s", err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, userInfo)
}
