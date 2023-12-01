package password

import (
	"chatgpt_x/pkg/e"
	"crypto/rand"
	"golang.org/x/crypto/bcrypt"
	"math/big"
)

// Hash 使用 bcrypt 对密码进行加密
func Hash(password string) string {
	// GenerateFromPassword 的第二个参数是 cost 值。建议大于 12，数值越大，耗时约长，安全性越高。
	// 性能：cost 值每 +1 耗时呈指数上升
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 13)
	e.HasError(err)
	return string(bytes)
}

// CheckHash 对比明文密码和数据库的哈希值
func CheckHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	e.HasError(err)
	return err == nil
}

// IsHashed 判断字符串是否是哈希过的数据
func IsHashed(str string) bool {
	// bcrypt 后的长度等于 60
	return len(str) == 60
}

// GetRandomPassword 获取一个随机密码
func GetRandomPassword(n int, allowedChars ...[]rune) string {
	var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	var letters []rune
	if len(allowedChars) == 0 {
		letters = defaultLetters
	} else {
		letters = allowedChars[0]
	}
	b := make([]rune, n)
	for i := range b {
		p, _ := rand.Int(rand.Reader, big.NewInt(10))
		b[i] = letters[int(p.Int64())]
	}
	return string(b)
}
