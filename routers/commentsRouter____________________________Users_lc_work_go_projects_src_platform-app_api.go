package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["platform-app/api:AppVersionHandler"] = append(beego.GlobalControllerRouter["platform-app/api:AppVersionHandler"],
		beego.ControllerComments{
			Method: "Upload",
			Router: `/upload`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
