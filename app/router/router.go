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
		// gin 的路由没有覆盖机制，所以只能把路由加长
		// user	相关路由
		userApi := apiRouter.Group("/user")
		{
			userApi.POST("/add", api.AddUser) // 增
			userApi.DELETE("/delete/:id", api.DeleteUser) // 删
			userApi.PUT("/edit/:id", api.EditUser) // 改
			userApi.GET("/list", api.GetUserPage) // 列表
			userApi.GET("/one/:id", api.GetUserById) // id查
		}
		// 栏目相关
		channelApi := apiRouter.Group("/channel")
		{
			channelApi.POST("/add", api.AddChannel)
			channelApi.DELETE("/delete/:id", api.DeleteChannel)
			channelApi.PUT("/edit/:id", api.EditChannel)
			channelApi.GET("/list", api.GetChannelPage)
			channelApi.GET("/tree", api.GetChannelTree)
			channelApi.GET("/one/:id", api.GetChannelById)
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
