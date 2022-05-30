# 腾讯云轻量控制面板

## 计划功能

- 控制台基本功能
- 轻量服务器快照托管，定时增删
- dnspod + cdn 融合面板，类 Cloudflare 
- 流量超过阀值停机，或自动换服务器
- 网站状态监控，自动切换服务器
- 一键部署热门应用，基于Docker镜像
- 跨账号管理资源
- TAT 命令集

## 安装方法

1、创建一个自定义策略 `TDPCloudAccess`，权限JSON如下：

```
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

2、创建一个用户，并关联策略 `TDPCloudAccess`。
