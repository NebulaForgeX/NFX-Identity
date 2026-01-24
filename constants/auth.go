package constants

// 认证相关常量
const (
	// MaxLoginAttempts 最大登录失败次数，超过此次数将锁定账户
	MaxLoginAttempts = 5

	// LockoutDurationMinutes 账户锁定持续时间（分钟）
	LockoutDurationMinutes = 30

	// DefaultAccessTokenTTLSeconds 默认 Access Token 有效期（秒），15 分钟
	DefaultAccessTokenTTLSeconds = 900

	// DefaultRefreshTokenTTLSeconds 默认 Refresh Token 有效期（秒），7 天
	DefaultRefreshTokenTTLSeconds = 7 * 24 * 3600
)
