# 土豆片控制面板

[![TDP Cloud Builder](https://github.com/open-tdp/tdp-cloud/actions/workflows/release.yml/badge.svg)](https://github.com/open-tdp/tdp-cloud/actions/workflows/release.yml)

可以跨平台部署的云资源管理面板

- 支持管理多个云账号资源

  - **腾讯云**（含国际版）：DNSPod、CVM、Lighthouse
  - **阿里云**（含国际版）：AliDNS (*doing*)、ECS、SWAS
  - **CloudFlare**：DNS、Custom Hostnames

- 支持添加子节点 (`TDP Worker`)

  - **Linux**：完整功能
  - **Macos**：部分功能
  - **Windows**：部分功能
  - **Android**：有限支持

- 支持自动签发基于`ACME`协议的SSL证书

  - **Let's Encrypt**：泛域名90天、单域名90天
  - **Bubpass**：泛域名90天、单域名180天
  - **Googel Public**：泛域名90天、单域名90天
  - **SSL.com**：泛域名90天、单域名90天
  - **ZeroSSL**：泛域名90天、单域名90天

## 使用指引

内容较多，请参考文档 <https://apps.rehiy.com/tdp-cloud/docs>

## 功能预览

- 在线体验开发版功能，请进入 [演示站点](https://apps.rehiy.com/tdp-cloud/preview)

- 功能支持和开发进度，请参阅 [Issues #1](https://github.com/open-tdp/tdp-cloud/issues/1)

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
