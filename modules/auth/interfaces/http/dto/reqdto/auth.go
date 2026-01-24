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

// SendVerificationCodeRequestDTO 发送验证码请求
type SendVerificationCodeRequestDTO struct {
	Email string `json:"email" validate:"required,email"`
}

// SignupRequestDTO 注册请求
type SignupRequestDTO struct {
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=6"`
	VerificationCode string `json:"verification_code" validate:"required,len=6"`
}
