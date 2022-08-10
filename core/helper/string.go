package helper

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"strings"
	"time"
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
