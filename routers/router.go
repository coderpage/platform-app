package routers

import (
	"platform-app/api"
	"platform-app/app"

	"github.com/astaxie/beego"
)

func init() {
	apiNameSpace := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/version",
			beego.NSInclude(
				&api.AppVersionHandler{})))
	beego.Include(&app.DownloadHandler{})
	beego.AddNamespace(apiNameSpace)
}

// Register for package import
func Register() {}
