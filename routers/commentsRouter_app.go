package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["platform-app/app:DownloadHandler"] = append(beego.GlobalControllerRouter["platform-app/app:DownloadHandler"],
		beego.ControllerComments{
			Method: "Download",
			Router: `/apk/:fileName`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
