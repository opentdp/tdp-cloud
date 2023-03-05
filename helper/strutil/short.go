package strutil

import (
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
)

// 转换编码 gb18030 -> utf-8

func Gb18030ToUtf8(s string) string {

	ret, err := simplifiedchinese.GB18030.NewDecoder().String(s)

	if err == nil {
		return string(ret)
	}
	return s

}

// 字符串首字母大写

func FirstUpper(s string) string {

	if s == "" {
		return s
	}

	return strings.ToUpper(s[:1]) + s[1:]

}

// 字符串首字母小写

func FirstLower(s string) string {

	if s == "" {
		return s
	}

	return strings.ToLower(s[:1]) + s[1:]

}
