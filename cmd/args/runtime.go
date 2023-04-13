package args

import (
	"embed"
)

// 嵌入文件

var Efs *embed.FS

// 子命令参数

var SubCommand struct {
	Name   string
	Action string
}
