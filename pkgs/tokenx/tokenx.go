package tokenx

// Tokenx 统一的 Token 管理器（类似 AWS Cognito）
type Tokenx struct {
	generator *Generator
	verifier  *Verifier
	cfg       Config
}

// New 创建新的 Tokenx 实例
func New(cfg Config) *Tokenx {
	return &Tokenx{
		generator: NewGenerator(cfg),
		verifier:  NewVerifier(cfg),
		cfg:       cfg,
	}
}

// GenerateAccessToken 生成 Access Token
func (t *Tokenx) GenerateAccessToken(userID, username, email, phone, roleID string) (string, error) {
	return t.generator.GenerateAccessToken(userID, username, email, phone, roleID)
}

// GenerateRefreshToken 生成 Refresh Token
func (t *Tokenx) GenerateRefreshToken(userID, username, email, phone, roleID string) (string, error) {
	return t.generator.GenerateRefreshToken(userID, username, email, phone, roleID)
}

// GenerateTokenPair 生成 Token 对
func (t *Tokenx) GenerateTokenPair(userID, username, email, phone, roleID string) (accessToken, refreshToken string, err error) {
	return t.generator.GenerateTokenPair(userID, username, email, phone, roleID)
}

// VerifyAccessToken 验证 Access Token
func (t *Tokenx) VerifyAccessToken(tokenString string) (*TokenClaims, error) {
	return t.verifier.VerifyAccessToken(tokenString)
}

// VerifyRefreshToken 验证 Refresh Token
func (t *Tokenx) VerifyRefreshToken(tokenString string) (*TokenClaims, error) {
	return t.verifier.VerifyRefreshToken(tokenString)
}

// VerifyToken 验证 Token（不检查类型）
func (t *Tokenx) VerifyToken(tokenString string) (*TokenClaims, error) {
	return t.verifier.VerifyToken(tokenString)
}

// RefreshTokenPair 刷新 Token 对
func (t *Tokenx) RefreshTokenPair(refreshToken string) (accessToken, newRefreshToken string, err error) {
	return t.verifier.RefreshTokenPair(refreshToken, t.generator)
}
