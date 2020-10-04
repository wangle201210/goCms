package router

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/wangle201210/goCms/app/middleware"
	"github.com/wangle201210/goCms/app/router/api"
	"github.com/wangle201210/goCms/app/util"
)

func Start() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	apiRouter := r.Group("/api")
	apiRouter.POST("/token", api.GetAuth)
	apiRouter.Use(middleware.AUTH())
	{
		// 获取当前登录用户信息
		apiRouter.GET("/mine", api.MineInfo)
		userRouter := apiRouter.Group("/user")
		{
			userRouter.GET("/show/:id", api.GetUserById)
			userRouter.GET("/list", api.GetUserPage)
			userRouter.POST("/add", api.AddUser)
		}

	}
	return r
}

func init() {
	gin.SetMode(util.ServerSetting.RunMode)
	r := Start()
	addr := fmt.Sprintf(":%d", util.ServerSetting.HttpPort)
	r.Run(addr)
}
