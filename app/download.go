package app

import (
	"net/http"
	"platform-app/controller"
	"platform-app/model"

	"github.com/astaxie/beego/orm"
)

// AppVersionHandler app verison controller
type DownloadHandler struct {
	controller.BaseController
}

//Download 下载
//@router /apk/:fileName [get]
func (handler *DownloadHandler) Download() {
	token := handler.GetString("token", "")
	if token == "" {
		handler.CustomAbort(http.StatusForbidden, "缺少参数 token")
	}

	appVersion, err := model.FindVersionByToken(token)
	if err != nil {
		if err == orm.ErrNoRows {
			handler.CustomAbort(http.StatusForbidden, "令牌错误")
		}
		handler.CustomAbort(http.StatusInternalServerError, err.Error())
	}

	model.IncreaseDownloadCount(appVersion.Id)

	fileName := handler.Ctx.Input.Param(":fileName")
	handler.Ctx.Output.Download("upload/apk/" + fileName)
}
