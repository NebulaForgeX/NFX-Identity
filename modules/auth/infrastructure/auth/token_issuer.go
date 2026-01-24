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
func (t *TokenIssuer) IssuePair(userID, username, email, phone, countryCode, roleID string) (access, refresh string, err error) {
	return t.tx.GenerateTokenPair(userID, username, email, phone, countryCode, roleID)
}

// IssuePairWithRefreshID 实现 TokenIssuer（带 refresh token ID）
func (t *TokenIssuer) IssuePairWithRefreshID(userID, username, email, phone, countryCode, roleID, refreshTokenID string) (access, refresh string, err error) {
	return t.tx.GenerateTokenPairWithRefreshID(userID, username, email, phone, countryCode, roleID, refreshTokenID)
}

// RefreshPair 实现 TokenIssuer
func (t *TokenIssuer) RefreshPair(refreshToken string) (access, refresh string, err error) {
	return t.tx.RefreshTokenPair(refreshToken)
}

// VerifyRefreshToken 验证 refresh token 并返回 claims
func (t *TokenIssuer) VerifyRefreshToken(refreshToken string) (*authApp.TokenClaims, error) {
	claims, err := t.tx.VerifyRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}
	
	tokenID := ""
	if claims.ID != "" {
		tokenID = claims.ID
	}
	
	return &authApp.TokenClaims{
		TokenID:     tokenID,
		UserID:      claims.UserID,
		Username:    claims.Username,
		Email:       claims.Email,
		Phone:       claims.Phone,
		CountryCode: claims.CountryCode,
		RoleID:      claims.RoleID,
	}, nil
}

var _ authApp.TokenIssuer = (*TokenIssuer)(nil)
