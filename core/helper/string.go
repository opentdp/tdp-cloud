package helper

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
)

// 序列化

func ToJsonString(r any) string {

	b, _ := json.Marshal(r)

	return string(b)

}

// 随机字符串

func RandString(length uint) string {

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

// 根据编码转换 byte 为 string

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

func Byte2String(byte []byte, charset Charset) string {

	var str string

	switch charset {
	case GB18030:
		ret, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(ret)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}

	return str

}
