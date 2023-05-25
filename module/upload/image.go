package upload

import (
	"encoding/base64"
	"os"
	"path"
	"time"

	"github.com/open-tdp/go-helper/strutil"

	"tdp-cloud/cmd/args"
)

func TimePathname(rand uint) string {

	base := time.Now().Format("/2006/0102/1504/05")

	if rand > 0 {
		base += strutil.Rand(rand)
	}

	return base

}

func SaveBase64Image(filePath, base64Image string) error {

	filePath = args.Dataset.Dir + "/upload" + filePath
	os.MkdirAll(path.Dir(filePath), 0755) // 递归创建目录

	imageBytes, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(imageBytes)
	if err != nil {
		return err
	}

	return nil

}
