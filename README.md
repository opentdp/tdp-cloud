
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
