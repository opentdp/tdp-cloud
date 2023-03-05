package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(str string) []byte {

	h := md5.New()
	h.Write([]byte(str))
	return h.Sum(nil)

}

func Md5ToString(str string) string {

	return hex.EncodeToString(Md5(str))

}
