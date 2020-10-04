package api

import (
	"strconv"

	"github.com/astaxie/beego/pkg/infrastructure/logs"
	"github.com/gin-gonic/gin"

	"github.com/wangle201210/goCms/app/model"
)

func GetUserById(c *gin.Context) ()  {
	s, b := c.GetQuery("id")
	if !b {
		logs.Error("通过id查询时必须传入id")
		return
	}
	id, _ := strconv.Atoi(s)
	u := &model.User{}
	u.ID = id
	if err := u.GetById(); err != nil {
		logs.Warn("get user by id err: %s",err)
	}
	c.JSON(200,u)
	return
}
