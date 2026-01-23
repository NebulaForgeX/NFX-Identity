package commands

// LoginCmd 登录命令
type LoginCmd struct {
	LoginType   string  // "email" | "phone"
	Email       *string
	Phone       *string
	CountryCode *string
	Password    string
	IP          *string // 客户端 IP 地址（用于记录登录尝试）
}

// RefreshCmd 刷新 Token 命令
type RefreshCmd struct {
	RefreshToken string
}
