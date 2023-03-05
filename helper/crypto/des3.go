package crypto

import (
	"crypto/des"
)

func Des3CBCEncrypt(src, key, iv []byte) ([]byte, error) {

	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}

	return CBCEncrypt(block, src, iv)

}

func Des3CBCDecrypt(src, key, iv []byte) ([]byte, error) {

	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}

	return CBCDecrypt(block, src, iv)

}
