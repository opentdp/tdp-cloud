# 轻量控制面板

基于腾讯云API实现的轻量级云资源控制面板

## 一、功能列表

请参阅 https://github.com/tdp-resource/tdp-cloud/issues/1

## 二、开发说明

### 启动开发服务

在项目目录运行  `serve.bat` 或 `./serve.sh`

### 编译为二进制

在项目目录运行  `build.bat` 或 `./build.sh`，或者[下载二进制可执行程序](https://github.com/tdp-resource/tdp-cloud/releases)

### 额外参数设置

如果项目无法运行或编译，请尝试设置系统环境变量或临时环境变量

```shell
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

## 三、部署说明

1、运行编译好的二进制文件，此时会生成 `cloud.db` 数据库文件，请注意权限

2、浏览器打开 `http://localhost:7800`，注册一个账号

3、登录刚注册的账号，添加腾讯云访问密钥

### 支持的数据库

- **SQLite** 数据源格式 `cloud.db`

- **MySQL**  数据源格式 `user:password@tcp(localhost:3306)/dbname?charset=utf8&parseTime=True&loc=Local`

### 添加腾讯云密钥

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

3、进入 `TDP Cloud` 后台，密钥管理，添加获取到的 `SecretId` 和 `SecretKey`

## 四、License

[GPL-3.0](https://opensource.org/licenses/GPL-3.0)

Copyright (c) 2022 TDP Cloud
