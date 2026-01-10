package auth

import (
	"nfxid/events"
)

// AccountLockoutsInvalidateCacheEvent AccountLockouts 缓存清除事件
// EventType 会自动从类型名生成，TopicKey 通过嵌入 events.AuthTopic 自动提供
// Cache key 格式: {prefix[:namespace]}:entity:{id}
type AccountLockoutsInvalidateCacheEvent struct {
	events.AuthTopic
	ID        string `json:"id"`         // 要清除的 Account Lockout ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "account_lockout"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// LoginAttemptsInvalidateCacheEvent LoginAttempts 缓存清除事件
type LoginAttemptsInvalidateCacheEvent struct {
	events.AuthTopic
	ID        string `json:"id"`         // 要清除的 Login Attempt ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "login_attempt"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// MFAFactorsInvalidateCacheEvent MFAFactors 缓存清除事件
type MFAFactorsInvalidateCacheEvent struct {
	events.AuthTopic
	ID        string `json:"id"`         // 要清除的 MFA Factor ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "mfa_factor"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// PasswordHistoryInvalidateCacheEvent PasswordHistory 缓存清除事件
type PasswordHistoryInvalidateCacheEvent struct {
	events.AuthTopic
	ID        string `json:"id"`         // 要清除的 Password History ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "password_history"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// PasswordResetsInvalidateCacheEvent PasswordResets 缓存清除事件
type PasswordResetsInvalidateCacheEvent struct {
	events.AuthTopic
	ID        string `json:"id"`         // 要清除的 Password Reset ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "password_reset"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// RefreshTokensInvalidateCacheEvent RefreshTokens 缓存清除事件
type RefreshTokensInvalidateCacheEvent struct {
	events.AuthTopic
	ID        string `json:"id"`         // 要清除的 Refresh Token ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "refresh_token"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// SessionsInvalidateCacheEvent Sessions 缓存清除事件
type SessionsInvalidateCacheEvent struct {
	events.AuthTopic
	ID        string `json:"id"`         // 要清除的 Session ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "session"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// TrustedDevicesInvalidateCacheEvent TrustedDevices 缓存清除事件
type TrustedDevicesInvalidateCacheEvent struct {
	events.AuthTopic
	ID        string `json:"id"`         // 要清除的 Trusted Device ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "trusted_device"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// UserCredentialsInvalidateCacheEvent UserCredentials 缓存清除事件
type UserCredentialsInvalidateCacheEvent struct {
	events.AuthTopic
	ID        string `json:"id"`         // 要清除的 User Credential ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "user_credential"
	Namespace string `json:"namespace"` // Cache namespace，可选
}
