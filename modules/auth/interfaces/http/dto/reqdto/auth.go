package reqdto

// LoginRequestDTO 登录请求（邮箱或手机号 + 密码）
type LoginRequestDTO struct {
	LoginType   string  `json:"login_type" validate:"required,oneof=email phone"`
	Email       *string `json:"email,omitempty"`
	Phone       *string `json:"phone,omitempty"`
	CountryCode *string `json:"country_code,omitempty"`
	Password    string  `json:"password" validate:"required"`
}

// RefreshRequestDTO 刷新 Token 请求
type RefreshRequestDTO struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
