package main

import (
	"platform-app/model"
	"platform-app/routers"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	model.Register()
	routers.Register()
}

func main() {
	beego.Run()
}
