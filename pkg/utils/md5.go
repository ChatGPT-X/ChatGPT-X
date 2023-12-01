package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// Get32MD5Encode 获取一个32位md5加密后的字符串.
func Get32MD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// Get16MD5Encode 获取一个16位md5加密后的字符串.
func Get16MD5Encode(data string) string {
	return Get32MD5Encode(data)[8:24]
}
