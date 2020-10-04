package model

import "github.com/jinzhu/gorm"

type Channel struct {
	Base

	Title       string `json:"title"`
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
}

var ChannelType = []string{
	"article",
	"linker",
}
// 增
func (m *Channel) Add() (err error) {
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
func (m *Channel) GetById(id int) (data *Channel, err error) {
	err = db.Model(m).Where("id = ? and deleted_at is null", id).First(&data).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return
}

// 分页数据
func (m *Channel) GetPage(pageNum, pageSize int, maps interface{}) (data []*Channel, err error) {
	err = db.Model(m).Where(maps).Offset(pageNum).Limit(pageSize).Find(&data).Error
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



