package reqdto

// LoginByEmailRequestDTO 邮箱登录请求
type LoginByEmailRequestDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// LoginByPhoneRequestDTO 手机号登录请求
type LoginByPhoneRequestDTO struct {
	CountryCode string `json:"country_code" validate:"required"`
	Phone       string `json:"phone" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

// RefreshRequestDTO 刷新 Token 请求
type RefreshRequestDTO struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
