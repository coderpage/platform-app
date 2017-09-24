package model

import "github.com/astaxie/beego/orm"

// AppVersion 版本信息
type AppVersion struct {
	ID             int64  `json:"id"`           // id
	AppName        string `json:"appName"`      // 应用名称
	PackageName    string `json:"packageName"`  // 报名
	ChangeLog      string `json:"changeLog"`    // change log
	VersionCode    int64  `json:"versionCode"`  // version code
	VesionName     string `json:"versionName"`  // 版本号
	DownloadURL    string `json:"downloadUrl"`  // 下载地址
	FileName       string `json:"fileName"`     // 文件名称
	FileSize       int64  `json:"fileSize"`     // 文件大小
	IsRelease      bool   `json:"isRelease"`    // 是否 release
	Uploader       string `json:"uploader"`     // 上传人名称
	UploadDate     int64  `json:"uploadDate"`   // 上传时间
	UploaderAvatar string `json:"uploadAvatar"` // 上传人的头像地址
}

// AddNewVersion 添加新的版本
func AddNewVersion(appVersion *AppVersion) error {
	o := orm.NewOrm()
	_, err := o.Insert(appVersion)
	return err
}

// FindLatestVersion 通过包名查找最新的版本信息
func FindLatestVersion(packageName string) (*AppVersion, error) {
	appVersion := &AppVersion{PackageName: packageName}
	o := orm.NewOrm()
	err := o.QueryTable("AppVersion").Filter("PackageName", packageName).OrderBy("-UploadDate").Limit(1).One(appVersion)
	return appVersion, err
}
