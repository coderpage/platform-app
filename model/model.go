package model

import (
	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterModel(new(AppVersion))
}

// 注册数据库表
func Register() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	mysqlUser := beego.AppConfig.String("mysqlDbUser")
	mysqlDb := beego.AppConfig.String("mysqlDbName")
	mysqlPwd := beego.AppConfig.String("mysqlDbPass")

	//tcp(112.80.45.162:9099)
	orm.RegisterDataBase("default", "mysql", mysqlUser+":"+mysqlPwd+"@/"+mysqlDb+"?charset=utf8&loc=Local")

	// 开启 ORM 调试模式
	orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", false, true)
}
