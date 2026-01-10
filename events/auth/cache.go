package auth

import (
	"nfxid/events"
)

// AccountLockoutsInvalidateCacheEvent AccountLockouts 缓存清除事件
// EventType 会自动从类型名生成，TopicKey 通过嵌入 events.AuthTopic 自动提供
type AccountLockoutsInvalidateCacheEvent struct {
	events.AuthTopic
	ID string `json:"id"` // 要清除的 Account Lockout ID
}

// LoginAttemptsInvalidateCacheEvent LoginAttempts 缓存清除事件
type LoginAttemptsInvalidateCacheEvent struct {
	events.AuthTopic
	ID string `json:"id"` // 要清除的 Login Attempt ID
}

// MFAFactorsInvalidateCacheEvent MFAFactors 缓存清除事件
type MFAFactorsInvalidateCacheEvent struct {
	events.AuthTopic
	ID string `json:"id"` // 要清除的 MFA Factor ID
}

// PasswordHistoryInvalidateCacheEvent PasswordHistory 缓存清除事件
type PasswordHistoryInvalidateCacheEvent struct {
	events.AuthTopic
	ID string `json:"id"` // 要清除的 Password History ID
}

// PasswordResetsInvalidateCacheEvent PasswordResets 缓存清除事件
type PasswordResetsInvalidateCacheEvent struct {
	events.AuthTopic
	ID string `json:"id"` // 要清除的 Password Reset ID
}

// RefreshTokensInvalidateCacheEvent RefreshTokens 缓存清除事件
type RefreshTokensInvalidateCacheEvent struct {
	events.AuthTopic
	ID string `json:"id"` // 要清除的 Refresh Token ID
}

// SessionsInvalidateCacheEvent Sessions 缓存清除事件
type SessionsInvalidateCacheEvent struct {
	events.AuthTopic
	ID string `json:"id"` // 要清除的 Session ID
}

// TrustedDevicesInvalidateCacheEvent TrustedDevices 缓存清除事件
type TrustedDevicesInvalidateCacheEvent struct {
	events.AuthTopic
	ID string `json:"id"` // 要清除的 Trusted Device ID
}

// UserCredentialsInvalidateCacheEvent UserCredentials 缓存清除事件
type UserCredentialsInvalidateCacheEvent struct {
	events.AuthTopic
	ID string `json:"id"` // 要清除的 User Credential ID
}
