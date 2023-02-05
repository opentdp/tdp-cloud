package strutil

import (
	"crypto/md5"
	"fmt"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
)

// 计算 Md5 值

func Md5(s string) string {

	return fmt.Sprintf("%x", md5.Sum([]byte(s)))

}

// 转换编码 gb18030 -> utf-8

func Gb18030ToUtf8(s string) string {

	ret, _ := simplifiedchinese.GB18030.NewDecoder().String(s)
	return string(ret)

}

// 字符串首字母大写

func FirstUpper(s string) string {

	return strings.ToUpper(s[:1]) + s[1:]

}

// 字符串首字母小写

func FirstLower(s string) string {

	return strings.ToLower(s[:1]) + s[1:]

}
