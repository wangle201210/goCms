package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/auth/jwt"

	"github.com/wangle201210/goCms/app/util"
)

// 用户登录验证的中间键
func AUTH() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := util.SUCCESS
		token := c.GetHeader("Authorization")
		userInfo, err := util.ParseToken(token[7:])
		if err != nil {
			switch err {
			case jwt.ErrTokenExpired:
				// 超时
				code = util.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			default:
				// 验证没通过
				code = util.ERROR_AUTH_CHECK_TOKEN_FAIL
			}
		}
		if code != util.SUCCESS {
			c.JSON(http.StatusUnauthorized,gin.H{
				"code":code,
				"msg": util.ErrMsg(code),
				"data": userInfo,
			})
			c.Abort()
			return
		}
		c.Set("userInfo",userInfo)
		//fmt.Println(userInfo)
		c.Next()
	}
}
