package upload

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"github.com/astaxie/beego/pkg/infrastructure/logs"

	"github.com/wangle201210/goCms/app/util"
	"github.com/wangle201210/goCms/app/util/file"
)

// 获取完整的图片路由
func GetImageFullUrl(name string) string {
	return util.AppSetting.ImagePrefixUrl + "/" + GetImagePath() + name
}

// 取文件内容为文件名，避免同文件多次传
func GetImageName(file multipart.File, image string) (string, error) {
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	ext := path.Ext(image)
	fileName := util.EncodeMD5(string(content))

	return fileName + ext, nil
}

func GetImagePath() string {
	return util.AppSetting.ImageSavePath
}

func GetImageFullPath() string {
	return util.AppSetting.FileUploadPath + GetImagePath()
}

func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range util.AppSetting.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

// 校验图片大小
func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		logs.Warn(err)
		return false
	}
	return size <= util.AppSetting.ImageMaxSize
}

func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
