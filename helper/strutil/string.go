package strutil

import (
	"crypto/md5"
	"fmt"

	"golang.org/x/text/encoding/simplifiedchinese"
)

// 计算 Md5 值

func Md5(s string) string {

	ret := md5.Sum([]byte(s))
	return fmt.Sprintf("%x", ret)

}

// 转换编码 gb18030 -> utf-8

func Gb18030ToUtf8(s string) string {

	ret, _ := simplifiedchinese.GB18030.NewDecoder().String(s)

	return string(ret)

}
