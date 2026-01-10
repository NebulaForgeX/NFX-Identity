package clients

import (
	"nfxid/events"
)

// APIKeysInvalidateCacheEvent APIKeys 缓存清除事件
// EventType 会自动从类型名生成，TopicKey 通过嵌入 events.ClientsTopic 自动提供
type APIKeysInvalidateCacheEvent struct {
	events.ClientsTopic
	ID string `json:"id"` // 要清除的 API Key ID
}

// AppsInvalidateCacheEvent Apps 缓存清除事件
type AppsInvalidateCacheEvent struct {
	events.ClientsTopic
	ID string `json:"id"` // 要清除的 App ID
}

// ClientCredentialsInvalidateCacheEvent ClientCredentials 缓存清除事件
type ClientCredentialsInvalidateCacheEvent struct {
	events.ClientsTopic
	ID string `json:"id"` // 要清除的 Client Credential ID
}

// ClientScopesInvalidateCacheEvent ClientScopes 缓存清除事件
type ClientScopesInvalidateCacheEvent struct {
	events.ClientsTopic
	ID string `json:"id"` // 要清除的 Client Scope ID
}

// IPAllowlistInvalidateCacheEvent IPAllowlist 缓存清除事件
type IPAllowlistInvalidateCacheEvent struct {
	events.ClientsTopic
	ID string `json:"id"` // 要清除的 IP Allowlist ID
}

// RateLimitsInvalidateCacheEvent RateLimits 缓存清除事件
type RateLimitsInvalidateCacheEvent struct {
	events.ClientsTopic
	ID string `json:"id"` // 要清除的 Rate Limit ID
}
