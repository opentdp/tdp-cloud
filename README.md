# 土豆片控制面板

基于腾讯云API实现的轻量级云资源控制面板

## 一、功能列表

支持的功能和开发进度，请参阅 [Issues #1](https://github.com/tdp-resource/tdp-cloud/issues/1)

WebUI界面请查看文档 [界面预览](https://github.com/tdp-resource/tdp-cloud/blob/main/docs/界面预览.md)

## 二、开发说明

### 启动开发服务

在项目目录运行  `serve.bat` 或 `./serve.sh`

### 编译为二进制

在项目目录运行 `build.bat` 或 `./build.sh`。你还可以下载 [稳定版](https://github.com/tdp-resource/tdp-cloud/releases) 或 [午夜构建版](http://curl.rpc.im/?dir=/tdp-cloud)

### 额外参数设置

如果项目无法运行或编译，请尝试设置系统环境变量或临时环境变量

```shell
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

## 三、服务端部署说明

1、根据系统类型下载编译好的[二进制程序](https://github.com/tdp-resource/tdp-cloud/releases)，重命名为 `tdpc`

2、运行 `tdpc server`，此时会生成 `server.db` 数据库文件，请注意权限

3、浏览器打开 `http://localhost:7800`，默认账号 `admin`，密码 `123456`

4、登录账号，添加一组或多组 `腾讯云CAM`，即可管理云资源

### 数据库配置参数

- 使用 **SQLite** 数据源 `--dsn "data/server.db"`。默认会追加参数 `?_pragma=busy_timeout=5000&_pragma=journa_mode(WAL)`

- 使用 **MySQL**  数据源 `--dsn "user:password@tcp(localhost:3306)/dbname"`。默认会追加参数 `?charset=utf8mb4&parseTime=True&loc=Local`

**注意：** 如果 `dsn` 字符串任意位置含有 `?` ，将忽略所有追加参数

### 添加腾讯云CAM

1、进入腾讯云 [CAM - 策略](https://console.cloud.tencent.com/cam/policy) 页面，创建一个自定义策略 `TDPCloudAccess`，权限JSON如下：

```json
{
    "version": "2.0",
    "statement": [
        {
            "action": [
                "cam:GetAccountSummary",
                "dnspod:*",
                "lighthouse:*",
                "monitor:*",
                "tat:*"
            ],
            "resource": "*",
            "effect": "allow"
        }
    ]
}
```

2、进入腾讯云 [CAM - 用户](https://console.cloud.tencent.com/cam) 页面，创建一个用户，允许 `编程访问`，并关联策略 `TDPCloudAccess`

3、进入 `TDP Cloud` 后台，厂商管理，添加获取到的 `SecretId` 和 `SecretKey`

## 四、其他

License [GPL-3.0](https://opensource.org/licenses/GPL-3.0)

Copyright (c) 2022 - 2023 TDP Cloud
