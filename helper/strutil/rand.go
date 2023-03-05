package strutil

import (
	"math/rand"
	"strconv"
	"strings"
)

// 随机字符串

func Rand(length uint) string {

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
