# 土豆片控制面板

[![TDP Cloud Builder](https://github.com/opentdp/tdp-cloud/actions/workflows/release.yml/badge.svg)](https://github.com/opentdp/tdp-cloud/actions/workflows/release.yml)

可以跨平台部署的云资源管理面板

- 支持管理多个云账号资源

  - **腾讯云**（*含国际版*）：DNSPod、CVM、Lighthouse
  - **阿里云**（*含国际版*）：AliDNS、ECS、SWAS
  - **CloudFlare**：DNS、Custom Hostnames

- 支持添加子节点 (`TDP Worker`)

  - **Linux**：完整功能
  - **Macos**：部分功能
  - **Windows**：部分功能
  - **Android**：有限支持

- 支持自动签发`SSL证书`，CA对比参见[使用指引](#使用指引)

  - **Let's Encrypt**
  - **Buypass**
  - **Googel Public**
  - **SSL.com**
  - **ZeroSSL**

- 支持 `WebSSH` 终端及`密钥对`管理

  - 支持使用已存储的密钥快速登录

  - 支持执行快捷命令

- 支持敏感数据加密存储（`3DES`）

  - 安装时，生成`通用密钥`，并保存至配置文件

  - 添加敏感资源时，将部分字段加密后存储至数据库

## 使用指引

内容较多，请参考文档 <https://docs.opentdp.org>

## 功能预览

- 功能支持和开发进度，请参阅 [Issues #1](https://github.com/opentdp/tdp-cloud/issues/1)

- 在线体验开发版功能，请进入 [演示站点](https://cloud.opentdp.org)，自行注册账号后登录

## 开发说明

### 初始化开发环境

```shell
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
go mod tidy
```

### 启动开发服务

在项目目录运行  `serve.bat` 或 `./serve.sh`

### 编译为二进制

在项目目录运行 `build.bat` 或 `./build.sh`。你还可以下载 [稳定版](https://cloud.opentdp.org/files)

### 提交代码时请使用下面标识

- `feat` 新功能（feature）
- `fix` 错误修复
- `docs` 文档更改（documentation）
- `style` 格式（不影响代码含义的更改，空格、格式、缺少分号等）
- `refactor` 重构（即不是新功能，也不是修补bug的代码变动）
- `perf` 优化（提高性能的代码更改）
- `test` 测试（添加缺失的测试或更正现有测试）
- `chore` 构建过程或辅助工具的变动
- `revert` 还原以前的提交

> 自 **v0.5.0** 起，`git commit` 描述请以 **标识+半角冒号+空格** 开头，即 `<type>: <subject>`

## 微信交流群

扫码添加开发者好友（请备注 `OpenTDP`，不备注可能无法通过好友申请）

![扫码加群](https://docs.opentdp.org/static/weixin-qr.jpg)

## 其他

License [GPL-3.0](https://www.gnu.org/licenses/gpl-3.0.txt)

Copyright (c) 2022 - 2023 OpenTDP
