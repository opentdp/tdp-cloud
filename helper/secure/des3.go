package secure

import (
	"encoding/hex"

	"github.com/forgoer/openssl"
)

// 加密字符串

func Des3Encrypt(s, p string) (string, error) {

	sb := []byte(s)

	pb := openssl.Md5ToString(p)
	secret, err := openssl.Des3CBCEncrypt(sb, []byte(pb[:24]), []byte(pb[24:]), openssl.PKCS7_PADDING)

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

	pb := openssl.Md5ToString(p)
	secret, err := openssl.Des3CBCDecrypt(sb, []byte(pb[:24]), []byte(pb[24:]), openssl.PKCS7_PADDING)

	if err == nil {
		return string(secret), err
	}
	return "", err

}
