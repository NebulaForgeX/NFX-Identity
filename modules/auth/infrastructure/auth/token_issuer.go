package auth

import (
	authApp "nfxid/modules/auth/application/auth"
	"nfxid/pkgs/tokenx"
)

// TokenIssuer 实现 application/auth.TokenIssuer，委托 tokenx
type TokenIssuer struct {
	tx *tokenx.Tokenx
}

// NewTokenIssuer 创建 TokenIssuer 适配器
func NewTokenIssuer(tx *tokenx.Tokenx) *TokenIssuer {
	return &TokenIssuer{tx: tx}
}

// IssuePair 实现 TokenIssuer
func (t *TokenIssuer) IssuePair(userID, username, email, phone, roleID string) (access, refresh string, err error) {
	return t.tx.GenerateTokenPair(userID, username, email, phone, roleID)
}

// RefreshPair 实现 TokenIssuer
func (t *TokenIssuer) RefreshPair(refreshToken string) (access, refresh string, err error) {
	return t.tx.RefreshTokenPair(refreshToken)
}

var _ authApp.TokenIssuer = (*TokenIssuer)(nil)
