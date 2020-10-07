package model

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/wangle201210/goCms/app/util"
)

// todo 媒体文件管理
type Media struct {
	Base

	OrgName   string `json:"org_name"`
	FileName  string `json:"file_name"`
	FilePath  string `json:"file_path"`
	Size      int    `json:"size"`
	ModelType string `json:"model_type"`
	ModelId   int    `json:"model_id"`
}

// 获取路由中允许被查询的字段
func (m *Media) GetQuery(c *gin.Context) (qm map[string]interface{}) {
	query := []string{
		"id",
		"OrgName",
		"file_name",
		"file_path",
		"model_type",
		"model_id",
	}
	res := make(map[string]interface{})
	for _, q := range query {
		if v, e := c.GetQuery(q); e {
			res[q] = v
		}
	}
	return res
}

// 增
func (m *Media) Add() (err error) {
	return db.Create(m).Error
}

// 删
func (m *Media) Delete() (err error) {
	return db.Delete(m).Error
}

// 改
func (m *Media) Edit(id int, data interface{}) (err error) {
	return db.Model(m).Where("id = ? and deleted_at is null", id).Updates(data).Error
}

// 通过id查找数据
func (m *Media) GetById() (err error) {
	err = db.Model(m).Where("id = ? and deleted_at is null", m.ID).First(&m).Error
	if err != nil {
		return err
	}
	return nil
}

// 分页数据
func (m *Media) GetPage(pageNum, maps interface{}) (data []*Media, err error) {
	err = db.Model(m).Where(maps).Offset(pageNum).Limit(util.AppSetting.PageSize).Find(&data).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return
}

// 获取总条数
func (m *Media) GetCount(maps interface{}) (count int, err error) {
	if err := db.Model(m).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}
	return
}

// 判断某条数据是否存在
func (m *Media) Exist() (exist bool, err error) {
	err = db.Select("id").Where("id = ? AND deleted_at is null", m.ID).First(m).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	if m.ID > 0 {
		exist = true
	}
	return
}
