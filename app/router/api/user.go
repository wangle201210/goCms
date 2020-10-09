package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/wangle201210/goCms/app/model"
	"github.com/wangle201210/goCms/app/util"
)

// 增
func AddUser(c *gin.Context) {
	g := util.Gin{C: c}
	u := &model.User{}
	if err := c.ShouldBind(&u); err != nil {
		g.Response(http.StatusBadRequest, util.INVALID_PARAMS, err.Error())
		return
	}
	u.Password = util.EncodeMD5(u.Password)
	if err := u.Add(); err != nil {
		g.Response(http.StatusBadRequest, util.ERROR_DATA_ADD, err.Error())
		return
	}
	g.Response(http.StatusOK, util.SUCCESS, u)
}

// 删
func DeleteUser(c *gin.Context)  {
	u := &model.User{}
	g := util.Gin{C: c}
	if !util.IsAdmin(c) {
		g.Response(http.StatusBadRequest, util.ERROR_AUTH_PERMISSION,nil)
		return
	}
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		g.Response(http.StatusBadRequest, util.INVALID_PARAMS,"id 必须为数字")
		return
	}
	if util.Mine(c).Id == id {
		g.Response(http.StatusBadRequest, util.INVALID_PARAMS,"不可删除自身账号")
		return
	}
	u.ID = id
	if err := u.Delete(); err != nil {
		g.Response(http.StatusBadRequest, util.ERROR_DATA_DELETE,err.Error())
		return
	}
	g.Response(http.StatusOK, util.SUCCESS,"删除成功！")
}

// 改
// todo 修改role时额外验证
func EditUser(c *gin.Context)  {
	g := util.Gin{C: c}
	s := c.Param("id")
	id, _ := strconv.Atoi(s)
	u := &model.User{}
	if err := c.ShouldBind(u); err != nil {
		g.Response(http.StatusBadRequest,util.INVALID_PARAMS,err.Error())
		return
	}
	u.ID = id
	if u.Password != "" {
		u.Password = util.EncodeMD5(u.Password)
	}
	if err := u.Edit(u.ID, u); err != nil {
		g.Response(http.StatusBadRequest,util.ERROR_DATA_EDIT,err.Error())
		return
	}
	g.Response(http.StatusOK,util.SUCCESS,nil)
}

func GetUserById(c *gin.Context) {
	g := util.Gin{C: c}
	s := c.Param("id")
	id, _ := strconv.Atoi(s)
	u := &model.User{}
	u.ID = id
	if err := u.GetById(); err != nil {
		g.Response(http.StatusBadRequest, util.ERROR_DATA_NOT_EXIST, nil)
		return
	}
	g.Response(http.StatusOK, util.SUCCESS, u)
}

func GetUserPage(c *gin.Context) {
	var (
		page int
		err  error
		g    util.Gin
	)
	g.C = c
	pageNum, exist := c.GetQuery("pageNum")

	if !exist {
		pageNum = "1"
	}
	page, err = strconv.Atoi(pageNum)
	if err != nil {
		g.Response(http.StatusBadRequest, util.INVALID_PARAMS, err.Error())
		return
	}
	start := util.GetPageStart(page)
	user := &model.User{}
	getPage, err := user.GetPage(start, user.GetQuery(c))
	if err != nil {
		g.Response(http.StatusBadRequest, util.INVALID_PARAMS, err.Error())
		return
	}
	total, err := user.GetCount(user.GetQuery(c))
	if err != nil {
		g.Response(http.StatusBadRequest, util.INVALID_PARAMS, err.Error())
		return
	}
	res := gin.H{
		"list":  getPage,
		"total": total,
	}
	g.Response(http.StatusOK, util.SUCCESS, res)
}
