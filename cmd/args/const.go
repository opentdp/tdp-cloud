package args

import (
	"embed"
)

var FrontFS *embed.FS

var Version = "0.5.1"
var BuildVersion = "202301"

var ConfigFile = "tdp-cloud.yml"

var ReadmeText = `土豆片控制面板

开源项目 https://github.com/tdp-resource/tdp-cloud
问题提交 https://github.com/tdp-resource/tdp-cloud/issues`
