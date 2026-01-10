package clients

import (
	"nfxid/events"
)

// APIKeysInvalidateCacheEvent APIKeys 缓存清除事件
// EventType 会自动从类型名生成，TopicKey 通过嵌入 events.ClientsTopic 自动提供
// Cache key 格式: {prefix[:namespace]}:entity:{id}
type APIKeysInvalidateCacheEvent struct {
	events.ClientsTopic
	ID        string `json:"id"`         // 要清除的 API Key ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "api_key"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// AppsInvalidateCacheEvent Apps 缓存清除事件
type AppsInvalidateCacheEvent struct {
	events.ClientsTopic
	ID        string `json:"id"`         // 要清除的 App ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "app"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// ClientCredentialsInvalidateCacheEvent ClientCredentials 缓存清除事件
type ClientCredentialsInvalidateCacheEvent struct {
	events.ClientsTopic
	ID        string `json:"id"`         // 要清除的 Client Credential ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "client_credential"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// ClientScopesInvalidateCacheEvent ClientScopes 缓存清除事件
type ClientScopesInvalidateCacheEvent struct {
	events.ClientsTopic
	ID        string `json:"id"`         // 要清除的 Client Scope ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "client_scope"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// IPAllowlistInvalidateCacheEvent IPAllowlist 缓存清除事件
type IPAllowlistInvalidateCacheEvent struct {
	events.ClientsTopic
	ID        string `json:"id"`         // 要清除的 IP Allowlist ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "ip_allowlist"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// RateLimitsInvalidateCacheEvent RateLimits 缓存清除事件
type RateLimitsInvalidateCacheEvent struct {
	events.ClientsTopic
	ID        string `json:"id"`         // 要清除的 Rate Limit ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "rate_limit"
	Namespace string `json:"namespace"` // Cache namespace，可选
}
