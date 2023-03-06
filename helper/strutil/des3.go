package strutil

import (
	"encoding/hex"

	"tdp-cloud/helper/crypto"
)

// 加密字符串

func Des3Encrypt(s, p string) (string, error) {

	sb := []byte(s)

	pb := crypto.Md5ToString(p)
	secret, err := crypto.Des3CBCEncrypt(sb, []byte(pb[:24]), []byte(pb[24:]))

	if err == nil {
		return hex.EncodeToString(secret), err
	}
	return "", err

}

// 解密字符串

func Des3Decrypt(s, p string) (string, error) {

	sb, err := hex.DecodeString(s)
	if err != nil {
		return "", err
	}

	pb := crypto.Md5ToString(p)
	secret, err := crypto.Des3CBCDecrypt(sb, []byte(pb[:24]), []byte(pb[24:]))

	if err == nil {
		return string(secret), err
	}
	return "", err

}
