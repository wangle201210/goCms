package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Mine(c *gin.Context) (info *Claims) {
	g := Gin{C: c}
	userInfo, exist := c.Get("userInfo")
	if !exist {
		g.Response(http.StatusUnauthorized, ERROR_AUTH_CHECK_TOKEN_FAIL, ErrMsg(ERROR_AUTH_CHECK_TOKEN_FAIL))
		c.Abort()
		return
	}
	info = userInfo.(*Claims)
	return
}

func IsAdmin(c *gin.Context) bool {
	mine := Mine(c)
	return mine.Role == 1
}

