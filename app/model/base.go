package model

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/pkg/infrastructure/logs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/wangle201210/goCms/app/util"
)

type Base struct {
	ID        int        `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

var db *gorm.DB

// 连接数据库
func init() {
	var err error
	db, err = gorm.Open(util.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		util.DatabaseSetting.User,
		util.DatabaseSetting.Password,
		util.DatabaseSetting.Host,
		util.DatabaseSetting.Port,
		util.DatabaseSetting.Name))

	if err != nil {
		logs.Error("connect database err: %v", err)
	}
	// 加上表前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return util.DatabaseSetting.TablePrefix + defaultTableName
	}

	db.SingularTable(true)

	// 部分配置
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	createTables()
	initData()
}

// 自动迁移
func createTables() {
	db.AutoMigrate(&User{}, &Article{}, &Channel{}, &Album{},&Linker{},&List{},&Media{})
}

//
//
//// 新增数据时自动添加 CreatedAt 和 UpdatedAt 的内容
//func updateTimeStampForCreateCallback(scope *gorm.Scope) {
//	if !scope.HasError() {
//		nowTime := time.Now()
//		if createTimeField, ok := scope.FieldByName("CreatedAt"); ok {
//			if createTimeField.IsBlank {
//				createTimeField.Set(nowTime)
//			}
//		}
//
//		if modifyTimeField, ok := scope.FieldByName("UpdatedAt"); ok {
//			if modifyTimeField.IsBlank {
//				modifyTimeField.Set(nowTime)
//			}
//		}
//	}
//}
//
//// 更新数据时修改 UpdatedAt 的内容
//func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
//	if _, ok := scope.Get("gorm:update_column"); !ok {
//		scope.SetColumn("UpdatedAt", time.Now())
//	}
//}
//
//// 删除数据时只更新 DeletedAt 做软删除
//func deleteCallback(scope *gorm.Scope) {
//	if !scope.HasError() {
//		var extraOption string
//		if str, ok := scope.Get("gorm:delete_option"); ok {
//			extraOption = fmt.Sprint(str)
//		}
//
//		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedAt")
//
//		if !scope.Search.Unscoped && hasDeletedOnField {
//			scope.Raw(fmt.Sprintf(
//				"UPDATE %v SET %v=%v %v %v",
//				scope.QuotedTableName(),
//				scope.Quote(deletedOnField.DBName),
//				scope.AddToVars(time.Now()),
//				scope.CombinedConditionSql(),
//				extraOption,
//			)).Exec()
//		} else {
//			scope.Raw(fmt.Sprintf(
//				"DELETE FROM %v %v %v",
//				scope.QuotedTableName(),
//				scope.CombinedConditionSql(),
//				extraOption,
//			)).Exec()
//		}
//	}
//}
