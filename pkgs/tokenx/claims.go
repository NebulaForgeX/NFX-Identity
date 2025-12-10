package tokenx

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// TokenClaims JWT Token Claims
type TokenClaims struct {
	UserID   string   `json:"user_id"`   // 用户ID
	Username string   `json:"username"`  // 用户名
	Email    string   `json:"email"`     // 邮箱
	Phone    string   `json:"phone"`     // 手机号
	RoleID   string   `json:"role_id"`   // 角色ID
	Type     TokenType `json:"type"`     // Token 类型：access 或 refresh
	jwt.RegisteredClaims
}

// TokenType Token 类型
type TokenType string

const (
	TokenTypeAccess  TokenType = "access"  // Access Token
	TokenTypeRefresh TokenType = "refresh" // Refresh Token
)

// NewAccessTokenClaims 创建 Access Token Claims
func NewAccessTokenClaims(userID, username, email, phone, roleID, issuer string, ttl time.Duration) *TokenClaims {
	now := time.Now()
	return &TokenClaims{
		UserID:   userID,
		Username: username,
		Email:    email,
		Phone:    phone,
		RoleID:   roleID,
		Type:     TokenTypeAccess,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,
			Subject:   userID,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
			NotBefore: jwt.NewNumericDate(now),
		},
	}
}

// NewRefreshTokenClaims 创建 Refresh Token Claims
func NewRefreshTokenClaims(userID, username, email, phone, roleID, issuer string, ttl time.Duration) *TokenClaims {
	now := time.Now()
	return &TokenClaims{
		UserID:   userID,
		Username: username,
		Email:    email,
		Phone:    phone,
		RoleID:   roleID,
		Type:     TokenTypeRefresh,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,
			Subject:   userID,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
			NotBefore: jwt.NewNumericDate(now),
		},
	}
}
