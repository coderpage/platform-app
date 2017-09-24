package app

import "platform-app/controller"

// AppVersionHandler app verison controller
type DownloadHandler struct {
	controller.BaseController
}

//Download 下载
//@router /apk/:fileName [get]
func (handler *DownloadHandler) Download() {
	fileName := handler.Ctx.Input.Param(":fileName")
	handler.Ctx.Output.Download("upload/apk/" + fileName)
}
