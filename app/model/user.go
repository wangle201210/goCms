package model

import (
	"github.com/astaxie/beego/pkg/infrastructure/logs"
	"github.com/jinzhu/gorm"

	"github.com/wangle201210/goCms/app/util"
)

type User struct {
	Base

	Name     string `json:"name"`
	Password string `json:"-"`
	Role     int    `json:"role"`
	Email    string `json:"email"`
}

var RoleList = []string{
	"unKnown",
	"admin",
}
// 增
func (m *User) Add() (err error) {
	return db.Create(m).Error
}

// 删
func (m *User) Delete() (err error) {
	return db.Delete(m).Error
}

// 改
func (m *User) Edit(id int, data interface{}) (err error) {
	// Update 一次只能更新一个字段
	// Updates 可以通过map更新多个字段
	return db.Model(m).Where("id = ? and deleted_at is null", id).Updates(data).Error
}

// 通过id查找数据
func (m *User) GetById() (err error) {
	err = db.Model(m).Where("id = ? and deleted_at is null", m.ID).First(&m).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	return nil
}

// 通过name查找数据
func (m *User) GetByName() (err error) {
	err = db.Model(m).Where("name = ? and deleted_at is null", m.Name).First(&m).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	return nil
}

// 分页数据
func (m *User) GetPage(pageNum, pageSize int, maps interface{}) (data []*Channel, err error) {
	err = db.Model(m).Where(maps).Offset(pageNum).Limit(pageSize).Find(&data).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return
}

// 获取总条数
func (m *User) GetCount(maps interface{}) (count int, err error) {
	if err := db.Model(m).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}
	return
}

// 判断某条数据是否存在
func (m *User) Exist(exist bool, err error) {
	err = db.Select("id").Where("id = ? AND deleted_at is null", m.ID).First(m).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	if m.ID > 0 {
		exist = true
	}
	return
}

// 添加初始用户
func initData() {
	user := &User{}
	err := db.First(&user).Error
	if err == gorm.ErrRecordNotFound {
		user.Name = "admin"
		user.Email = "iwangle.me@gmail.com"
		user.Role = 1
		user.Password = util.EncodeMD5("password")
		err = user.Add()
	}
	if err != nil {
		logs.Error("query err: %s", err)
	}
}
