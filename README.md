# 土豆片控制面板

[![TDP Cloud Builder](https://github.com/tdp-resource/tdp-cloud/actions/workflows/release.yml/badge.svg)](https://github.com/tdp-resource/tdp-cloud/actions/workflows/release.yml)

可以跨平台部署的云资源管理面板，支持同时绑定多个云账号，目前已实现管理下列云资源：

 - CloudFlare：DNS

 - 腾讯云：DNSPod、CVM、Lighthouse

 - 腾讯云国际版：CVM、Lighthouse

 - 独立主机：TDP Worker

##  功能预览

- 在线体验开发版功能，请进入 [演示站点](https://apps.rehiy.com/tdp-cloud/preview)

- 功能支持和开发进度，请参阅 [Issues #1](https://github.com/tdp-resource/tdp-cloud/issues/1)

## 开发说明

### 启动开发服务

在项目目录运行  `serve.bat` 或 `./serve.sh`

### 提交代码时请使用下面标识

- `feat` 新功能（feature）
- `fix` 修补bug
- `docs` 文档（documentation）
- `style` 格式（不影响代码运行的变动）
- `refactor` 重构（即不是新功能，也不是修补bug的代码变动）
- `perf` 优化
- `test` 增加测试
- `chore` 构建过程或辅助工具的变动
- `revert` 撤销某次操作

> 此项要求自 **v0.5.0** 起实施，所有 `git commit` 描述需要使用 **标识+半角冒号+空格** 开头，即 `<type>: <subject>`

### 编译为二进制

在项目目录运行 `build.bat` 或 `./build.sh`。你还可以下载 [稳定版](https://apps.rehiy.com/tdp-cloud/release) 或 [午夜构建版](https://apps.rehiy.com/tdp-cloud/nightly)

### 额外参数设置

如果项目无法运行或编译，请尝试设置系统环境变量或临时环境变量

```shell
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

## 使用说明

https://apps.rehiy.com/tdp-cloud/docs

## 其他

License [GPL-3.0](https://opensource.org/licenses/GPL-3.0)

Copyright (c) 2022 - 2023 TDP Cloud
