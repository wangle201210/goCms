package util

import (
	"time"

	"github.com/astaxie/beego/pkg/infrastructure/logs"
	"github.com/go-ini/ini"
)

type App struct {
	PageSize	int
	ExpireTime	int
	Name        string
	Md5String	string
	JwtSecret	string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Port        int
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string

}

var DatabaseSetting = &Database{}

// 设置项转map
// https://ini.unknwon.io/docs/advanced/map_and_reflect
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		logs.Error("Cfg.MapTo %s err: %v", section, err)
	}
}

var cfg *ini.File

func init() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		logs.Error("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}
	mapTo("app", AppSetting)
	mapTo("database", DatabaseSetting)
	mapTo("server", ServerSetting)
}
