# 腾讯云轻量控制面板

## 功能及进度

- [x] 注册/登录
- [ ] 控制台首页
- [x] 域名列表
- [x] 轻量服务器列表
- [ ] 轻量服务器快照托管，定时增删
- [ ] 网站状态监控，宕机自动换解析
- [ ] 流量超过阀值，自动关机、换解析
- [ ] 管理SSL证书
- [ ] Dnspod + CDN 融合面板，类 Cloudflare
- [ ] 一键部署热门应用，可能会基于 Docker 镜像
- [ ] 实现自定义 TAT 命令集
- [x] 跨账号管理资源，绑定多个CAM密钥

## 运行开发服务

在项目目录运行  `serve.bat` 或 `./serve.sh`

## 编译为二进制

在项目目录运行  `build.bat` 或 `./build.sh`

## 额外参数设置

如果项目无法运行或编译，请尝试设置系统环境变量或临时环境变量
  
```shell
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

## 初次部署说明

1、运行编译好的二进制文件，此时会生成 `cloud.db` 数据库文件，请注意权限

2、浏览器打开 `http://localhost:7800`，注册一个账号

3、登录刚注册的账号，添加腾讯云访问密钥

## 获取腾讯云密钥

1、创建一个自定义策略 `TDPCloudAccess`，权限JSON如下：

```json
{
    "version": "2.0",
    "statement": [
        {
            "action": [
                "cam:GetAccountSummary",
                "dnspod:*",
                "lighthouse:*"
            ],
            "resource": "*",
            "effect": "allow"
        }
    ]
}
```

2、创建一个用户，允许 `编程访问`，并关联策略 `TDPCloudAccess`

3、使用生成的 `SecretId` 和 `SecretId` 登录面板

# License

[GPL-3.0](https://opensource.org/licenses/GPL-3.0)

Copyright (c) 2022 TDP Cloud
