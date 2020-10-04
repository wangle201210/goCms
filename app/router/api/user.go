package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/wangle201210/goCms/app/model"
	"github.com/wangle201210/goCms/app/util"
)

func AddUser(c *gin.Context) {
	g := util.Gin{C: c}
	u := &model.User{}
	if err := c.ShouldBind(&u); err != nil {
		g.Response(http.StatusBadRequest, util.INVALID_PARAMS, err.Error())
		return
	}
	u.Password = util.EncodeMD5(u.Password)
	if err := u.Add(); err != nil {
		g.Response(http.StatusBadRequest, util.ERROR_USER_ADD_FAIL, err.Error())
		return
	}
	g.Response(http.StatusOK, util.SUCCESS, u)
	return
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
	return
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
	return
}
