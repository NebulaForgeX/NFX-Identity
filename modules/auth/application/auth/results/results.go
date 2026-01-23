package results

// LoginResult 登录结果
type LoginResult struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int64
	UserID       string
}

// RefreshResult 刷新 Token 结果
type RefreshResult struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int64
}
