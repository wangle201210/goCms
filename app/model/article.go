package model

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/wangle201210/goCms/app/util"
)

type Article struct {
	Base

	ChannelID   int    `json:"channel_id" gorm:"index"`
	Title       string `json:"title"`
	Keyword     string `json:"keyword"`
	Content     string `json:"content"`
	Description string `json:"description"`
	Pic         string `json:"pic"`
	Count       string `json:"count"`

	Channel *Channel `json:"channel" gorm:"-"`
}

// 获取路由中允许被查询的字段
func (m *Article) GetQuery(c *gin.Context) (qm map[string]interface{}) {
	query := []string{
		"id",
		"title",
		"keyword",
		"channel_id",
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
func (m *Article) Add() (err error) {
	return db.Create(m).Error
}

// 删
func (m *Article) Delete() (err error) {
	return db.Delete(m).Error
}

// 改
func (m *Article) Edit(id int, data interface{}) (err error) {
	return db.Model(m).Where("id = ? and deleted_at is null", id).Updates(data).Error
}

// 通过id查找数据
func (m *Article) GetById() (err error) {
	err = db.Model(m).Where("id = ? and deleted_at is null", m.ID).First(&m).Error
	if err != nil {
		return err
	}
	return nil
}

// 分页数据
func (m *Article) GetPage(pageNum, maps interface{}) (data []*Article, err error) {
	err = db.Model(m).Where(maps).Offset(pageNum).Limit(util.AppSetting.PageSize).Find(&data).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return
}

// 获取总条数
func (m *Article) GetCount(maps interface{}) (count int, err error) {
	if err := db.Model(m).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}
	return
}

// 判断某条数据是否存在
func (m *Article) Exist() (exist bool, err error) {
	err = db.Select("id").Where("id = ? AND deleted_at is null", m.ID).First(m).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	if m.ID > 0 {
		exist = true
	}
	return
}
