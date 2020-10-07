package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/wangle201210/goCms/app/model"
	"github.com/wangle201210/goCms/app/util"
)

type Auth struct {
	Name     string `json:"name" from:"name"`
	Password string `json:"password" from:"password"`
	Role     int    `json:"role" from:"role"`
}

func GetAuth(c *gin.Context) {
	g := util.Gin{C: c}
	auth := Auth{}
	if err := c.ShouldBind(&auth); err != nil {
		g.Response(http.StatusBadRequest, util.INVALID_PARAMS, nil)
		return
	}
	u := model.User{Name: auth.Name}
	if err := u.GetByName(); err != nil {
		// 用户不存在
		g.Response(http.StatusBadRequest, util.ERROR_AUTH_NOUSER, nil)
		return
	}
	// 用户存在则验证密码
	if util.EncodeMD5(auth.Password) != u.Password {
		g.Response(http.StatusBadRequest, util.ERROR_AUTH_PASSWORD, nil)
		return
	}
	token, err := util.GenerateToken(u.ID, u.Name, u.Role)
	if err != nil {
		g.Response(http.StatusBadRequest, util.ERROR_AUTH_TOKEN, nil)
		return
	}

	// 上面生成正常，这里解析就一定正常
	parseToken, _ := util.ParseToken(token)
	expiresAt := time.Unix(parseToken.ExpiresAt, 0)
	g.Response(http.StatusOK, util.SUCCESS, gin.H{
		"token":     token,
		"expiresAt": expiresAt.Format(util.AppSetting.TimeFormat),
	})
}

func MineInfo(c *gin.Context) {
	g := util.Gin{C: c}
	u := &model.User{}
	mine := util.Mine(c)
	u.ID = mine.Id
	if err := u.GetById(); err != nil {
		g.Response(http.StatusBadRequest, util.ERROR_DATA_NOT_EXIST, err.Error())
		return
	}
	g.Response(http.StatusOK, util.SUCCESS, u)
}

// todo 退出时考虑要不要处理什么逻辑
func Logout(c *gin.Context) {
	g := util.Gin{C: c}
	g.Response(http.StatusOK, util.SUCCESS, nil)
}
