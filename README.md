# 土豆片控制面板

已支持管理的云资源：

 - CloudFlare： DNS

 - 腾讯云： Lighthouse、CVM、DNSPod

 - 独立服务器：TDP Worker

##  功能列表

支持的功能和开发进度，请参阅 [Issues #1](https://github.com/tdp-resource/tdp-cloud/issues/1)

前端界面展示，请参看 [界面预览](https://github.com/tdp-resource/tdp-cloud/blob/main/docs/界面预览.md)

## 开发说明

### 启动开发服务

在项目目录运行  `serve.bat` 或 `./serve.sh`

### 编译为二进制

在项目目录运行 `build.bat` 或 `./build.sh`。你还可以下载 [稳定版](https://apps.rehiy.com/tdp-cloud/release) 或 [午夜构建版](https://apps.rehiy.com/tdp-cloud/nightly)

### 额外参数设置

如果项目无法运行或编译，请尝试设置系统环境变量或临时环境变量

```shell
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

## 部署说明

### 安装服务端

https://github.com/tdp-resource/tdp-cloud/blob/main/docs/部署服务端.md

### 添加子节点

https://github.com/tdp-resource/tdp-cloud/blob/main/docs/添加子节点.md

### 绑定腾讯云

https://github.com/tdp-resource/tdp-cloud/blob/main/docs/绑定腾讯云.md

### 绑定Cloudflare

https://github.com/tdp-resource/tdp-cloud/blob/main/docs/绑定Cloudflare.md

## 其他

License [GPL-3.0](https://opensource.org/licenses/GPL-3.0)

Copyright (c) 2022 - 2023 TDP Cloud
