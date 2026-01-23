package constants

// 认证相关常量
const (
	// MaxLoginAttempts 最大登录失败次数，超过此次数将锁定账户
	MaxLoginAttempts = 5

	// LockoutDurationMinutes 账户锁定持续时间（分钟）
	LockoutDurationMinutes = 30
)
