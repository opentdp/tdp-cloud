package args

import (
	"os"
	"path"
	"path/filepath"

	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/strutil"
)

// 初始化调试模式

func SetDebug() {

	de := os.Getenv("TDP_DEBUG")
	Debug = de == "1" || de == "true"

}

// 初始化存储目录

func SetAssets() {

	if Assets.Secret == "" {
		Assets.Secret = strutil.Rand(32)
	}

	if Assets.Dir != "" && Assets.Dir != "." {
		os.MkdirAll(Assets.Dir, 0755)
	}

}

// 初始化日志能力

func SetLogger() {

	storage := Logger.Dir
	if !filepath.IsAbs(storage) {
		storage = path.Join(Assets.Dir, storage)
	}

	logman.SetDefault(&logman.Config{
		Level:    Logger.Level,
		Target:   Logger.Target,
		Storage:  storage,
		Filename: "server",
	})

}
