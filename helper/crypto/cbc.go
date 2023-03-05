package crypto

import (
	"crypto/cipher"
	"errors"
)

func CBCEncrypt(block cipher.Block, src, iv []byte) ([]byte, error) {

	blockSize := block.BlockSize()
	src = PKCS7Padding(src, blockSize)

	encryptData := make([]byte, len(src))

	if len(iv) != block.BlockSize() {
		return nil, errors.New("CBCEncrypt: IV length must equal block size")
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(encryptData, src)

	return encryptData, nil

}

func CBCDecrypt(block cipher.Block, src, iv []byte) ([]byte, error) {

	dst := make([]byte, len(src))

	if len(iv) != block.BlockSize() {
		return nil, errors.New("CBCDecrypt: IV length must equal block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(dst, src)

	return PKCS7UnPadding(dst)

}
