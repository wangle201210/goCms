package model

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/wangle201210/goCms/app/util"
)

type Channel struct {
	Base

	Title       string `json:"title" gorm:"unique"`
	Keyword     string `json:"keyword"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Deep        int    `json:"deep"`
	ParentId    int    `json:"parent_id"`
	Show        int    `json:"show"`
	Target      int    `json:"target"`
	Pic         string `json:"pic"`
	OutLinker   int    `json:"out_linker"`
	RedirectUrl string `json:"redirect_url"`
	Ordering    int    `json:"ordering"`
	Count       int    `json:"count"`
	Style       string `json:"style"`

	Children []*Channel `json:"children" gorm: "-"'`
}

var ChannelType = []string{
	"article",
	"linker",
	"list",
}

// 获取路由中允许被查询的字段
// todo 想办法通过 Channel 反射得到字段列表
func (m *Channel) GetQuery(c *gin.Context) (qm map[string]interface{}) {
	query := []string{
		"id",
		"title",
		"keyword",
		"type",
		"parent_id",
		"show",
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
// todo title 唯一性
func (m *Channel) Add() (err error) {
	if err = m.GetByTitle(); err != nil {
		return
	} else if m.ID > 0 {
		err = errors.New(util.ErrMsg(util.ERROR_CHANNEL_TITLE_USED))
		return
	}

	return db.Create(m).Error
}

// 删
func (m *Channel) Delete() (err error) {
	return db.Delete(m).Error
}

// 改
func (m *Channel) Edit(id int, data interface{}) (err error) {
	// Update 一次只能更新一个字段
	// Updates 可以通过map更新多个字段
	return db.Model(m).Where("id = ? and deleted_at is null", id).Updates(data).Error
}

// 通过id查找数据
func (m *Channel) GetById() (err error) {
	err = db.Model(m).Where("id = ? and deleted_at is null", m.ID).First(&m).Error
	if err != nil {
		return err
	}
	return nil
}

// 通过name查找数据
func (m *Channel) GetByTitle() (err error) {
	err = db.Model(m).Where("title = ? and deleted_at is null", m.Title).First(&m).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	return nil
}

// 分页数据
func (m *Channel) GetPage(pageNum, maps interface{}) (data []*Channel, err error) {
	err = db.Model(m).Where(maps).Offset(pageNum).Limit(util.AppSetting.PageSize).Find(&data).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return
}

// 获取总条数
func (m *Channel) GetCount(maps interface{}) (count int, err error) {
	if err := db.Model(m).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}
	return
}

// 判断某条数据是否存在
func (m *Channel) Exist(exist bool, err error) {
	err = db.Select("id").Where("id = ? AND deleted_at is null", m.ID).First(m).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	if m.ID > 0 {
		exist = true
	}
	return
}

// 获取全部频道
func (m *Channel) GetAll(maps ...interface{}) (data []*Channel, err error) {
	if len(maps) > 0 {
		err = db.Model(m).Where(maps[0]).Find(&data).Error
	} else {
		err = db.Model(m).Find(&data).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return
}

func (m *Channel) MakeTree(data []*Channel, id, deep int) (res []*Channel, err error) {
	for _, v := range data {
		if v.ParentId != id {
			if p, e := m.findP(data, v.ParentId); e {
				p.Children = append(p.Children,v)
			}
		} else {
			res = append(res,v)
		}
	}
	return
}

func  (m *Channel) findP(data []*Channel, id int) (*Channel, bool) {
	for _, i2 := range data {
		if i2.ID == id {
			return i2, true
		}
	}
	return nil,false
}
