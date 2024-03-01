package fsadmin

import (
	"encoding/base64"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/opentdp/go-helper/strutil"

	"tdp-cloud/cmd/args"
)

const UploadDir = "/upload"

func TimePathname(rand uint) string {

	name := time.Now().Format("/2006/0102/1504/05")

	if rand > 0 {
		name += strutil.Rand(rand)
	}

	return name

}

func UintPathname(id uint) string {

	uid := strconv.FormatUint(uint64(id), 10)
	for len(uid) < 12 {
		uid = fmt.Sprintf("%012s", uid)
	}

	name := "/"
	name += uid[0:4] + "/"
	name += uid[4:8] + "/"
	name += uid[8:12] + "/"
	name += strconv.FormatInt(time.Now().Unix(), 10)

	return name

}

func SaveBase64Image(filePath, base64Image string) error {

	if !strings.HasPrefix(filePath, UploadDir) {
		filePath = UploadDir + filePath
	}

	filePath = path.Join(args.Assets.Dir, filePath)
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
