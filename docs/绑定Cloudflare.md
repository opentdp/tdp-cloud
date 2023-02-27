# 绑定Cloudflare

1、进入 **Cloudflare** [个人资料 - API 令牌](https://dash.cloudflare.com/profile/api-tokens) 页面，创建一个令牌 `TDPCloudAccess`，权限如下：

```text
User.User Details, Zone.Zone Settings, Zone.Zone, Zone.SSL and Certificates, Zone.DNS
```

3、进入 **TDP Cloud** 后台，`资产管理 - Cloudflare`，添加获取到的 `API Token`

4、在 `资产管理 - Cloudflare` 中选择刚添加的账号，点击 `导入` 按钮，选择需要导入资源，完成绑定操作
