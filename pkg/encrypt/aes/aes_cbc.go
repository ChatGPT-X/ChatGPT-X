package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// AesEncryptCBC AES-CBC加密.
func AesEncryptCBC(orig string, key []byte, iv []byte) string {
	origData := []byte(orig)

	block, _ := aes.NewCipher(key)

	blockSize := block.BlockSize()

	origData = PKCS7Padding(origData, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, iv)

	cryted := make([]byte, len(origData))

	blockMode.CryptBlocks(cryted, origData)
	return base64.StdEncoding.EncodeToString(cryted)
}

// AesDecryptCBC AES-CBC解密.
func AesDecryptCBC(cryted string, key []byte, iv []byte) string {
	crytedByte, _ := base64.StdEncoding.DecodeString(cryted)

	block, _ := aes.NewCipher(key)

	blockMode := cipher.NewCBCDecrypter(block, iv)

	orig := make([]byte, len(crytedByte))

	blockMode.CryptBlocks(orig, crytedByte)

	orig = PKCS7UnPadding(orig)
	return string(orig)
}
