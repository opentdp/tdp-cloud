# 土豆片控制面板

[![TDP Cloud Builder](https://github.com/open-tdp/tdp-cloud/actions/workflows/release.yml/badge.svg)](https://github.com/open-tdp/tdp-cloud/actions/workflows/release.yml)

可以跨平台部署的云资源管理面板

- 支持管理多个云账号资源

  - **腾讯云**（*含国际版*）：DNSPod、CVM、Lighthouse
  - **阿里云**（*含国际版*）：AliDNS (*doing*)、ECS、SWAS
  - **CloudFlare**：DNS、Custom Hostnames

- 支持添加子节点 (`TDP Worker`)

  - **Linux**：完整功能
  - **Macos**：部分功能
  - **Windows**：部分功能
  - **Android**：有限支持

- 支持自动签发`SSL证书`，CA对比参见[使用指引](#使用指引)

  - **Let's Encrypt**
  - **Bubpass**
  - **Googel Public**
  - **SSL.com**
  - **ZeroSSL**

- 支持 `WebSSH` 终端及`密钥对`管理

  - 支持使用已存储的密钥快速登录

  - 支持执行快捷命令

- 支持敏感数据加密存储（`3DES`）

  - 注册时，生成`AppId`和`AppKey`，并使用登陆密码加密`AppKey`后存储

  - 登陆时，使用登陆密码解密`AppKey`，并交由`JWT`保管

  - 新建厂商和密钥对时，使用`JWT`保管的`AppKey`加密私钥后存储

  - 调用资源时，使用`JWT`保管的`AppKey`还原私钥

## 使用指引

内容较多，请参考文档 <https://apps.rehiy.com/tdp-cloud/docs>

## 功能预览

- 功能支持和开发进度，请参阅 [Issues #1](https://github.com/open-tdp/tdp-cloud/issues/1)

- 在线体验开发版功能，请进入 [演示站点](https://apps.rehiy.com/tdp-cloud/preview)，自行注册账号后登录

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

## 其他

License [GPL-3.0](https://www.gnu.org/licenses/gpl-3.0.txt)

Copyright (c) 2022 - 2023 TDP Cloud
