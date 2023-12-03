package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

// jwtSecret 密钥，不得少于 64 字节。
var jwtSecret []byte

// jwtIssuer 签发人。
var jwtIssuer string

// jwtActiveTime 有效时间，单位：秒。
var jwtActiveTime int

// InitConfig 初始化配置。
func InitConfig(secret, issuer string, activeTime int) {
	jwtSecret = []byte(secret)
	jwtIssuer = issuer
	jwtActiveTime = activeTime
}

// CustomClaims 自定义声明。
type CustomClaims struct {
	UserID   uint   `json:"user_id"`
	IsAdmin  uint   `json:"is_admin"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Claims 声明。
type Claims struct {
	CustomClaims
	jwt.RegisteredClaims
}

// GenerateToken 生成 Token。
func GenerateToken(cc CustomClaims) (string, error) {
	// 计算过期时间
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(jwtActiveTime) * time.Second)
	// 构造 jwt 主结构
	claims := Claims{
		CustomClaims: cc,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    jwtIssuer,
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(nowTime),
		},
	}
	// 生成 Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString(jwtSecret)
}

// ParseToken 解析 Token。
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		return claims, nil
	}
	return nil, jwt.ErrSignatureInvalid
}

// CheckToken 检查 Token 是否合法。
func CheckToken(token string) bool {
	_, err := ParseToken(token)
	return err == nil
}

// GetTokenFromHeader 从 Header 中获取 Token。
func GetTokenFromHeader(c *gin.Context) (string, error) {
	token := c.GetHeader("Authorization")
	// token 为空或者不包含 Bearer  关键字
	if len(token) == 0 || !strings.Contains(token, "Bearer ") {
		return "", errors.New("jwt: token not found")
	}
	return strings.TrimPrefix(token, "Bearer "), nil
}
