package model

import (
	"errors"

	"github.com/astaxie/beego/pkg/infrastructure/logs"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/wangle201210/goCms/app/util"
)

type User struct {
	Base

	Name     string `json:"name" gorm:"not null;unique"` // 非空且唯一
	Password string `json:"password"`
	Role     int    `json:"role"`
	Email    string `json:"email"`
}

var RoleList = []string{
	"unKnown",
	"admin",
}

// 获取路由中允许被查询的字段
// todo 研究下为什么 map[string]interface{} 不能换成gin.H
func (m *User) GetQuery(c *gin.Context) (qm map[string]interface{}) {
	query := []string{
		"id",
		"name",
		"role",
		"email",
	}
	res := make(map[string]interface{})
	for _, q := range query {
		if v, e := c.GetQuery(q); e {
			res[q] = v
		}
	}
	return res
}

// 获取json中允许被修改（新增）的字段
//func (m *User) GetPost(c *gin.Context) (pm map[string]interface{}) {
//	post := []string{
//		"name",
//		"role",
//		"email",
//		"password",
//	}
//	res := make(map[string]interface{})
//	for _, q := range post {
//		if v, e := c.GetPostForm(q); e {
//			res[q] = v
//		}
//	}
//	return res
//}

// 增
// todo name 唯一性
func (m *User) Add() (err error) {
	if err = m.GetByName(); err != nil {
		return
	} else if m.ID > 0 {
		err = errors.New(util.ErrMsg(util.ERROR_USER_NAME_USED))
		return
	}

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
	if err != nil {
		return err
	}
	return nil
}

// 通过name查找数据
func (m *User) GetByName() (err error) {
	err = db.Model(m).Where("name = ? and deleted_at is null", m.Name).First(&m).Error
	if err != nil {
		return err
	}
	return nil
}

// 分页数据
func (m *User) GetPage(pageNum, maps interface{}) (data []*User, err error) {
	err = db.Model(m).Where(maps).Offset(pageNum).Limit(util.AppSetting.PageSize).Find(&data).Error
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
func (m *User) Exist() (exist bool, err error) {
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
