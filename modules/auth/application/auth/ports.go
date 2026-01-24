package auth

// TokenIssuer 签发与刷新 JWT。
// 由 infrastructure 实现（如 tokenx 适配器）。
type TokenIssuer interface {
	IssuePair(userID, username, email, phone, countryCode, roleID string) (access, refresh string, err error)
	IssuePairWithRefreshID(userID, username, email, phone, countryCode, roleID, refreshTokenID string) (access, refresh string, err error)
	RefreshPair(refreshToken string) (access, refresh string, err error)
	VerifyRefreshToken(refreshToken string) (*TokenClaims, error)
}

// TokenClaims 从 refresh token 中解析出的 claims（用于获取 token_id/jti）
type TokenClaims struct {
	TokenID     string // JWT 的 jti claim
	UserID      string
	Username    string
	Email       string
	Phone       string
	CountryCode string // 国家代码（手机号登录时使用）
	RoleID      string
}
