package respdto

// LoginResponseDTO 登录响应
type LoginResponseDTO struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"` // seconds
	UserID       string `json:"user_id"`
}

// RefreshResponseDTO 刷新 Token 响应
type RefreshResponseDTO struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"` // seconds
}
