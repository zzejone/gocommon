// Package jwtmanager
package jwtmanager

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

// CustomClaims 自定义Claims结构，使用map存储自定义数据
type CustomClaims struct {
	Data map[string]string `json:"data"` // 存储自定义数据
	jwt.RegisteredClaims
}

// JWTOptions JWT配置
type JWTOptions struct {
	SecretKey     string
	ExpiresAt     time.Duration
	Issuer        string
	SigningMethod jwt.SigningMethod
}

// JWTManager JWT管理器
type JWTManager struct {
	options *JWTOptions
}

type Option func(*JWTOptions)

// NewJWTManager 初始化JWT管理器
func NewJWTManager(opts ...Option) *JWTManager {
	options := &JWTOptions{
		SecretKey:     "oC5VCKMDsW!P",
		ExpiresAt:     2 * time.Hour,
		Issuer:        "jwtmanager",
		SigningMethod: jwt.SigningMethodHS256,
	}
	for _, opt := range opts {
		opt(options)
	}

	return &JWTManager{options: options}
}

func SetSecretKey(ipt string) Option {
	return func(j *JWTOptions) {
		j.SecretKey = ipt
	}
}

func SetExpiresAt(ipt time.Duration) Option {
	return func(j *JWTOptions) {
		j.ExpiresAt = ipt
	}
}

func SetIssuer(ipt string) Option {
	return func(j *JWTOptions) {
		j.Issuer = ipt
	}
}

func SetSignMethod(ipt jwt.SigningMethod) Option {
	return func(j *JWTOptions) {
		j.SigningMethod = ipt
	}
}

// GenerateToken 生成JWT Token - 支持map传入自定义数据
func (jm *JWTManager) GenerateToken(data map[string]string) (string, error) {
	// 设置Claims
	claims := CustomClaims{
		Data: data,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jm.options.ExpiresAt)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    jm.options.Issuer,
		},
	}

	// 创建Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名Token
	return token.SignedString([]byte(jm.options.SecretKey))
}

// ParseToken 解析JWT Token
func (jm *JWTManager) ParseToken(tokenString string) (map[string]string, error) {
	// 解析Token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (any, error) {
		// 验证签名方法是否匹配
		if token.Method != jm.options.SigningMethod {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jm.options.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	// 验证Token并获取Claims
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims.Data, nil
	}

	return nil, errors.New("invalid token")
}

// ParseTokenWithClaims 解析Token并返回完整的Claims信息（可选）
func (jm *JWTManager) ParseTokenWithClaims(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (any, error) {
		if token.Method != jm.options.SigningMethod {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jm.options.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// GetDataFromToken 从Token中获取指定字段的值
func (jm *JWTManager) GetDataFromToken(tokenString string, key string) (any, error) {
	data, err := jm.ParseToken(tokenString)
	if err != nil {
		return nil, err
	}

	if value, exists := data[key]; exists {
		return value, nil
	}

	return nil, errors.New("key not found in token data")
}

// GetStringFromToken 从Token中获取字符串值
func (jm *JWTManager) GetStringFromToken(tokenString string, key string) (string, error) {
	value, err := jm.GetDataFromToken(tokenString, key)
	if err != nil {
		return "", err
	}

	if str, ok := value.(string); ok {
		return str, nil
	}

	return "", errors.New("value is not a string")
}

// GetInt64FromToken 从Token中获取int64值
func (jm *JWTManager) GetInt64FromToken(tokenString string, key string) (int64, error) {
	value, err := jm.GetDataFromToken(tokenString, key)
	if err != nil {
		return 0, err
	}

	// 处理JSON数字类型（float64）
	if num, ok := value.(float64); ok {
		return int64(num), nil
	}

	if num, ok := value.(int64); ok {
		return num, nil
	}

	if num, ok := value.(int); ok {
		return int64(num), nil
	}

	return 0, errors.New("value is not a number")
}

// RefreshToken 刷新Token
func (jm *JWTManager) RefreshToken(tokenString string) (string, error) {
	data, err := jm.ParseToken(tokenString)
	if err != nil {
		return "", err
	}

	return jm.GenerateToken(data)
}
