package auth

// RegisterCmd 注册命令
type RegisterCmd struct {
	Email           string // 邮箱
	VerificationCode string // 验证码
	AuthorizationCode string // 授权码（邀请码）
	Password        string // 密码
}
