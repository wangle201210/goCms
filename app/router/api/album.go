package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/wangle201210/goCms/app/model"
	"github.com/wangle201210/goCms/app/util"
)

// 增
func AddAlbum(c *gin.Context) {
	g := util.Gin{C: c}
	m := &model.Album{}
	if err := c.ShouldBind(m); err != nil {
		g.Response(http.StatusBadRequest, util.INVALID_PARAMS, err.Error())
		return
	}
	if err := m.Add(); err != nil {
		g.Response(http.StatusBadRequest, util.ERROR_DATA_ADD, err.Error())
		return
	}
	g.Response(http.StatusOK, util.SUCCESS, nil)
}

// 删
func DeleteAlbum(c *gin.Context)  {
	g := util.Gin{C: c}
	m := &model.Album{}
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
	m.ID = id
	if err := m.Delete(); err != nil {
		g.Response(http.StatusBadRequest, util.ERROR_DATA_DELETE,err.Error())
		return
	}
	g.Response(http.StatusOK, util.SUCCESS,"删除成功！")
}

// 改
func EditAlbum(c *gin.Context)  {
	m := &model.Album{}
	g := util.Gin{C: c}
	s := c.Param("id")
	id, _ := strconv.Atoi(s)
	if err := c.ShouldBind(m); err != nil {
		g.Response(http.StatusBadRequest,util.INVALID_PARAMS,err.Error())
		return
	}
	m.ID = id
	if err := m.Edit(m.ID, m); err != nil {
		g.Response(http.StatusBadRequest,util.ERROR_DATA_EDIT,err.Error())
		return
	}
	g.Response(http.StatusOK,util.SUCCESS,nil)
}

func GetAlbumById(c *gin.Context) {
	m := &model.Album{}
	g := util.Gin{C: c}
	s := c.Param("id")
	id, _ := strconv.Atoi(s)
	m.ID = id
	if err := m.GetById(); err != nil {
		g.Response(http.StatusBadRequest, util.ERROR_DATA_NOT_EXIST, nil)
		return
	}
	g.Response(http.StatusOK, util.SUCCESS, m)
}

func GetAlbumPage(c *gin.Context) {
	var (
		page int
		err  error
		g    util.Gin
	)
	m := &model.Album{}
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
	getPage, err := m.GetPage(start, m.GetQuery(c))
	if err != nil {
		g.Response(http.StatusBadRequest, util.INVALID_PARAMS, err.Error())
		return
	}
	total, err := m.GetCount(m.GetQuery(c))
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
