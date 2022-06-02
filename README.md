# 腾讯云轻量控制面板

## 功能及进度

- [x] 使用子账号密钥登录
- [x] 控制台首页
  - [x] 域名列表
  - [x] 轻量服务器列表
- [ ] 轻量服务器快照托管，定时增删
- [ ] 网站状态监控，宕机自动换解析
- [ ] 流量超过阀值，自动关机、换解析
- [ ] Dnspod + CDN 融合面板，类 Cloudflare
- [ ] 一键部署热门应用，可能会基于 Docker 镜像
- [ ] 实现自定义 TAT 命令集
- [ ] 跨账号管理资源

## 运行开发服务

在项目目录运行  `serve.bat` 或 `./serve.sh`

## 编译为二进制

在项目目录运行  `build.bat` 或 `./build.sh`

## 额外参数设置

如果你的项目无法运行，请尝试如下设置：

```shell
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

## 获取登录密钥

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
