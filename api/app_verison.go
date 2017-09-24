package api

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"net/http"
	"os"
	"platform-app/controller"
	"platform-app/model"
	"strconv"
	"time"

	"github.com/lunny/axmlParser"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

const (
	APK_UPLOAD_DIR = "upload/apk/"
)

// AppVersionHandler app verison controller
type AppVersionHandler struct {
	controller.BaseController
}

type Size interface {
	Size() int64
}

// GetLatest 获取最新的版本
// @router /latest [get]
func (handler *AppVersionHandler) GetLatest() {
	packageName := handler.GetString("packageName", "")
	if packageName == "" {
		handler.CustomAbort(http.StatusBadRequest, "缺少参数 packageName")
	}

	appVersion, err := model.FindLatestVersion(packageName)

	if err != nil {
		if err == orm.ErrNoRows {
			handler.CustomAbort(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		}
		handler.CustomAbort(http.StatusInternalServerError, err.Error())
	}

	appVersion.DownloadURL = beego.AppConfig.String("serverBaseUrl") + "/apk/" + appVersion.FileName
	resp := handler.NewResponse()
	resp.SetStatus(http.StatusOK).SetData("data", appVersion)
	handler.Data["json"] = resp
	handler.ServeJSON()
}

// Upload 上传 apk 文件
// @router /upload [post]
func (handler *AppVersionHandler) Upload() {
	resp := handler.NewResponse()

	token := handler.GetString("token", "")
	appName := handler.GetString("appName", "")
	changeLog := handler.GetString("changeLog", "")
	isRelease, _ := handler.GetBool("isRelease", true)

	apkUploadTk := beego.AppConfig.String("apkUploadTk")
	if token != apkUploadTk {
		handler.CustomAbort(http.StatusForbidden, "令牌验证失败")
	}

	apkFile, fileHeader, err := handler.GetFile("file")
	if err != nil {
		handler.CustomAbort(http.StatusBadRequest, err.Error())
	}
	defer apkFile.Close()

	fileName := fileHeader.Filename
	sizeInterface, _ := apkFile.(Size)
	fileSize := sizeInterface.Size()

	md5Ctx := md5.New()
	_, err = io.Copy(md5Ctx, apkFile)
	if err == nil {
		// 使用 MD5 值作为文件名称
		fileName = hex.EncodeToString(md5Ctx.Sum([]byte(""))) + ".apk"
	} else {
		// MD5 计算失败，使用文件名+时间戳命名
		timetamp := strconv.FormatInt(time.Now().Unix(), 10)
		fileName = fileName + timetamp + ".apk"
	}

	logs.Debug("upload apk filename:", fileName, "size:", sizeInterface.Size())

	err = os.MkdirAll(APK_UPLOAD_DIR, 0700)

	if err != nil {
		logs.Error("create dir err:", err)
		resp.SetStatus(http.StatusInternalServerError).SetMessage(err.Error())
		handler.Data["json"] = resp
		handler.ServeJSON()
		return
	}

	err = handler.SaveToFile("file", APK_UPLOAD_DIR+fileName)
	if err != nil {
		logs.Error("save file err:", err)
		resp.SetStatus(http.StatusInternalServerError).SetMessage(err.Error())
		handler.Data["json"] = resp
		handler.ServeJSON()
		return
	}

	listener := new(axmlParser.AppNameListener)
	_, err = axmlParser.ParseApk(APK_UPLOAD_DIR+fileName, listener)
	if err != nil {
		logs.Error("parse apk file err:", err)
		resp.SetStatus(-1).SetMessage(err.Error())
		handler.Data["json"] = resp
		handler.ServeJSON()
		return
	}

	verisonCode, _ := strconv.ParseInt(listener.VersionCode, 10, 64)
	appVersion := &model.AppVersion{}
	appVersion.AppName = appName
	appVersion.ChangeLog = changeLog
	appVersion.PackageName = listener.PackageName
	appVersion.VersionCode = verisonCode
	appVersion.VesionName = listener.VersionName
	appVersion.FileSize = fileSize
	appVersion.FileName = fileName
	appVersion.DownloadURL = fileName
	appVersion.IsRelease = isRelease
	appVersion.UploadDate = time.Now().Unix()
	err = model.AddNewVersion(appVersion)
	if err != nil {
		resp.SetStatus(-1).SetMessage(err.Error())
		handler.Data["json"] = resp
		handler.ServeJSON()
		return
	}

	logs.Debug("APP VERSION", appVersion)

	appVersion.DownloadURL = beego.AppConfig.String("serverBaseUrl") + "/apk/" + appVersion.FileName

	resp.SetStatus(http.StatusOK).SetMessage(http.StatusText(http.StatusOK))
	resp.SetData("data", appVersion)
	handler.Data["json"] = resp
	handler.ServeJSON()
}
