package strings

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
)

// 随机字符串

func Rand(length uint) string {

	rand.Seed(time.Now().UnixNano())

	rs := make([]string, length)

	for i := uint(0); i < length; i++ {
		t := rand.Intn(3)
		if t == 0 {
			rs = append(rs, strconv.Itoa(rand.Intn(10)))
		} else if t == 1 {
			rs = append(rs, string(rune(rand.Intn(26)+65)))
		} else {
			rs = append(rs, string(rune(rand.Intn(26)+97)))
		}
	}

	return strings.Join(rs, "")

}

// 转换编码 gb18030 -> utf-8

func Gb18030ToUtf8(str string) string {

	ret, _ := simplifiedchinese.GB18030.NewDecoder().String(str)

	return string(ret)

}
