package commands

// LoginByEmailCmd 邮箱登录命令
type LoginByEmailCmd struct {
	Email    string
	Password string
	IP       *string // 客户端 IP 地址（用于记录登录尝试）
}

// LoginByPhoneCmd 手机号登录命令
type LoginByPhoneCmd struct {
	CountryCode string
	Phone       string
	Password    string
	IP          *string // 客户端 IP 地址（用于记录登录尝试）
}

// RefreshCmd 刷新 Token 命令
type RefreshCmd struct {
	RefreshToken string
}
