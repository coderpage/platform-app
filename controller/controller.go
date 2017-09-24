package controller

import (
	"encoding/json"
	"page/constant/rsp"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

// Http 请求返回数据
type Response map[string]interface{}

// 创建一个 Response
func (this *BaseController) NewResponse() (resp Response) {
	return make(Response)
}

// Http 返回数据的 json 格式字符串
func (resp Response) JsonString() string {
	respBytes, _ := json.Marshal(resp)
	return string(respBytes)
}

// 设置 response 的 status 值
func (resp Response) SetStatus(status interface{}) Response {
	resp[rsp.BodyStatus] = status
	return resp
}

// 设置 response 的 message 值
func (resp Response) SetMessage(message interface{}) Response {
	resp[rsp.BodyMsg] = message
	return resp
}

// 设置 response 的 data 值
func (resp Response) SetData(name string, data interface{}) {
	resp[name] = data
}
