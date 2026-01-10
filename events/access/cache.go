package access

import (
	"nfxid/events"
)

// GrantsInvalidateCacheEvent Grants 缓存清除事件
// EventType 会自动从类型名生成，TopicKey 通过嵌入 events.AccessTopic 自动提供
// Cache key 格式: {prefix[:namespace]}:entity:{id}
type GrantsInvalidateCacheEvent struct {
	events.AccessTopic
	ID        string `json:"id"`        // 要清除的 Grant ID
	Prefix    string `json:"prefix"`    // Cache prefix，例如 "grant"
	Namespace string `json:"namespace"` // Cache namespace，可选，例如 tenantId 或 service name
}

// PermissionsInvalidateCacheEvent Permissions 缓存清除事件
type PermissionsInvalidateCacheEvent struct {
	events.AccessTopic
	ID        string `json:"id"`        // 要清除的 Permission ID
	Prefix    string `json:"prefix"`    // Cache prefix，例如 "permission"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// RolePermissionsInvalidateCacheEvent RolePermissions 缓存清除事件
type RolePermissionsInvalidateCacheEvent struct {
	events.AccessTopic
	ID        string `json:"id"`        // 要清除的 RolePermission ID
	Prefix    string `json:"prefix"`    // Cache prefix，例如 "role_permission"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// RolesInvalidateCacheEvent Roles 缓存清除事件
type RolesInvalidateCacheEvent struct {
	events.AccessTopic
	ID        string `json:"id"`        // 要清除的 Role ID
	Prefix    string `json:"prefix"`    // Cache prefix，例如 "role"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// ScopePermissionsInvalidateCacheEvent ScopePermissions 缓存清除事件
type ScopePermissionsInvalidateCacheEvent struct {
	events.AccessTopic
	ID        string `json:"id"`        // 要清除的 ScopePermission ID
	Prefix    string `json:"prefix"`    // Cache prefix，例如 "scope_permission"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// ScopesInvalidateCacheEvent Scopes 缓存清除事件
type ScopesInvalidateCacheEvent struct {
	events.AccessTopic
	ID        string `json:"id"`        // 要清除的 Scope ID (通常是 scope 字符串)
	Prefix    string `json:"prefix"`    // Cache prefix，例如 "scope"
	Namespace string `json:"namespace"` // Cache namespace，可选
}
