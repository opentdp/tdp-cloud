package crypto

import (
	"bytes"
	"errors"
)

func PKCS7Padding(src []byte, blockSize int) []byte {

	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(src, padtext...)

}

func PKCS7UnPadding(src []byte) ([]byte, error) {

	length := len(src)
	if length == 0 {
		return src, errors.New("unpadding error")
	}

	unpadding := int(src[length-1])
	if length < unpadding {
		return src, errors.New("unpadding error")
	}

	return src[:(length - unpadding)], nil

}
