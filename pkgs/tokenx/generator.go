package tokenx

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

// Generator Token 生成器
type Generator struct {
	cfg Config
}

// NewGenerator 创建新的 Token 生成器
func NewGenerator(cfg Config) *Generator {
	return &Generator{
		cfg: cfg,
	}
}

// GenerateAccessToken 生成 Access Token
func (g *Generator) GenerateAccessToken(userID, username, email, phone, countryCode, roleID string) (string, error) {
	claims := NewAccessTokenClaims(userID, username, email, phone, countryCode, roleID, g.cfg.Issuer, g.cfg.AccessTokenTTL)
	return g.signToken(claims)
}

// GenerateRefreshToken 生成 Refresh Token
func (g *Generator) GenerateRefreshToken(userID, username, email, phone, countryCode, roleID string) (string, error) {
	claims := NewRefreshTokenClaims(userID, username, email, phone, countryCode, roleID, g.cfg.Issuer, g.cfg.RefreshTokenTTL)
	return g.signToken(claims)
}

// GenerateRefreshTokenWithID 生成 Refresh Token（带 token_id/jti）
func (g *Generator) GenerateRefreshTokenWithID(userID, username, email, phone, countryCode, roleID, tokenID string) (string, error) {
	claims := NewRefreshTokenClaimsWithID(userID, username, email, phone, countryCode, roleID, g.cfg.Issuer, g.cfg.RefreshTokenTTL, tokenID)
	return g.signToken(claims)
}

// GenerateTokenPair 生成 Token 对（Access + Refresh）
func (g *Generator) GenerateTokenPair(userID, username, email, phone, countryCode, roleID string) (accessToken, refreshToken string, err error) {
	accessToken, err = g.GenerateAccessToken(userID, username, email, phone, countryCode, roleID)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate access token: %w", err)
	}

	refreshToken, err = g.GenerateRefreshToken(userID, username, email, phone, countryCode, roleID)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate refresh token: %w", err)
	}

	return accessToken, refreshToken, nil
}

// GenerateTokenPairWithRefreshID 生成 Token 对（Access + Refresh，refresh token 带 token_id/jti）
func (g *Generator) GenerateTokenPairWithRefreshID(userID, username, email, phone, countryCode, roleID, refreshTokenID string) (accessToken, refreshToken string, err error) {
	accessToken, err = g.GenerateAccessToken(userID, username, email, phone, countryCode, roleID)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate access token: %w", err)
	}

	refreshToken, err = g.GenerateRefreshTokenWithID(userID, username, email, phone, countryCode, roleID, refreshTokenID)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate refresh token: %w", err)
	}

	return accessToken, refreshToken, nil
}

// signToken 签名 Token
func (g *Generator) signToken(claims *TokenClaims) (string, error) {
	method := jwt.GetSigningMethod(g.cfg.Algorithm)
	if method == nil {
		return "", fmt.Errorf("unsupported signing method: %s", g.cfg.Algorithm)
	}

	token := jwt.NewWithClaims(method, claims)
	return token.SignedString([]byte(g.cfg.SecretKey))
}
