# 腾讯云轻量控制面板

## 计划功能

- 自动管理快照
- dnspod + cdn 融合面板，参考 Cloudflare 
- 流量超过阀值停机，或自动换服务器
- 网站状态监控，自动切换服务器

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
