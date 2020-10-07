package api

import (
	"net/http"

	"github.com/astaxie/beego/pkg/infrastructure/logs"
	"github.com/gin-gonic/gin"

	"github.com/wangle201210/goCms/app/util"
	"github.com/wangle201210/goCms/app/util/upload"
)

func UploadImage(c *gin.Context) {
	g := util.Gin{C: c}

	file, image, err := c.Request.FormFile("image")
	if err != nil {
		g.Response(http.StatusBadRequest, util.INVALID_PARAMS, err.Error())
		return
	}

	if image == nil {
		g.Response(http.StatusBadRequest, util.INVALID_PARAMS, nil)
		return
	}
	imageName, err := upload.GetImageName(file, image.Filename)
	if err != nil {
		g.Response(http.StatusBadRequest, util.INVALID_PARAMS, err.Error())
		return
	}
	fullPath := upload.GetImageFullPath()
	savePath := upload.GetImagePath()

	src := fullPath + imageName
	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		g.Response(http.StatusBadRequest, util.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, nil)
		return
	}

	if err := upload.CheckImage(fullPath); err != nil {
		logs.Warn(err)
		g.Response(http.StatusBadRequest, util.ERROR_UPLOAD_CHECK_IMAGE_FAIL, err.Error())
		return
	} else if err := c.SaveUploadedFile(image, src); err != nil {
		logs.Warn(err)
		g.Response(http.StatusBadRequest, util.ERROR_UPLOAD_SAVE_IMAGE_FAIL, err.Error())
		return
	}

	data := gin.H{}
	data["image_url"] = upload.GetImageFullUrl(imageName)
	data["image_save_url"] = savePath + imageName
	g.Response(http.StatusOK, util.SUCCESS, data)
}
