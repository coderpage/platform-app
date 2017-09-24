# platform-app

应用上传升级 API 服务

1. 获取应用最新版本

```
METHOD : GET
URL : /api/v1/version/latest
QUERY : packageName

http://127.0.0.1:8001/api/v1/version/latest?packageName=com.coderpage.mine
```

2. 上传应用

```
METHOD : POST
URL : /api/v1/version/upload
MultipartBody:
  file : apk 文件 (file)
  token : 令牌 (string)
  appName : 应用名称 (string)
  changeLog : changeLog (string)
  isRelease : is release (bool)
```
