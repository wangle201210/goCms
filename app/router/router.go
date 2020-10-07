package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/wangle201210/goCms/app/middleware"
	"github.com/wangle201210/goCms/app/router/api"
	"github.com/wangle201210/goCms/app/util"
	"github.com/wangle201210/goCms/app/util/upload"
)

func Start() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CORS())
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	apiRouter := r.Group("/api")
	apiRouter.POST("/token", api.GetAuth)
	apiRouter.Use(middleware.AUTH())
	{
		apiRouter.POST("/logout", api.Logout)

		// 获取当前登录用户信息
		apiRouter.GET("/mine", api.MineInfo)
		// 图片上传
		apiRouter.POST("/upload/image", api.UploadImage)

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

		// 文章相关
		articleApi := apiRouter.Group("/article")
		{
			articleApi.POST("/add", api.AddArticle)
			articleApi.DELETE("/delete/:id", api.DeleteArticle)
			articleApi.PUT("/edit/:id", api.EditArticle)
			articleApi.GET("/list", api.GetArticlePage)
			articleApi.GET("/one/:id", api.GetArticleById)
		}

		// 新闻列表相关
		listApi := apiRouter.Group("/list")
		{
			listApi.POST("/add", api.AddList)
			listApi.DELETE("/delete/:id", api.DeleteList)
			listApi.PUT("/edit/:id", api.EditList)
			listApi.GET("/list", api.GetListPage)
			listApi.GET("/one/:id", api.GetListById)
		}

		// 友情链接相关
		linkerApi := apiRouter.Group("/linker")
		{
			linkerApi.POST("/add", api.AddLinker)
			linkerApi.DELETE("/delete/:id", api.DeleteLinker)
			linkerApi.PUT("/edit/:id", api.EditLinker)
			linkerApi.GET("/list", api.GetLinkerPage)
			linkerApi.GET("/one/:id", api.GetLinkerById)
		}

		// 相册相关
		albumApi := apiRouter.Group("/album")
		{
			albumApi.POST("/add", api.AddAlbum)
			albumApi.DELETE("/delete/:id", api.DeleteAlbum)
			albumApi.PUT("/edit/:id", api.EditAlbum)
			albumApi.GET("/list", api.GetAlbumPage)
			albumApi.GET("/one/:id", api.GetAlbumById)
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
