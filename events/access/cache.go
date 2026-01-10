package access

import (
	"nfxid/events"
)

// GrantsInvalidateCacheEvent Grants 缓存清除事件
// EventType 会自动从类型名生成，TopicKey 通过嵌入 events.AccessTopic 自动提供
type GrantsInvalidateCacheEvent struct {
	events.AccessTopic
	ID          string `json:"id"`           // 要清除的 Grant ID
	ServiceName string `json:"service_name"` // 服务名称
}

// PermissionsInvalidateCacheEvent Permissions 缓存清除事件
type PermissionsInvalidateCacheEvent struct {
	events.AccessTopic
	ID string `json:"id"` // 要清除的 Permission ID
}

// RolePermissionsInvalidateCacheEvent RolePermissions 缓存清除事件
type RolePermissionsInvalidateCacheEvent struct {
	events.AccessTopic
	ID string `json:"id"` // 要清除的 RolePermission ID
}

// RolesInvalidateCacheEvent Roles 缓存清除事件
type RolesInvalidateCacheEvent struct {
	events.AccessTopic
	ID string `json:"id"` // 要清除的 Role ID
}

// ScopePermissionsInvalidateCacheEvent ScopePermissions 缓存清除事件
type ScopePermissionsInvalidateCacheEvent struct {
	events.AccessTopic
	ID string `json:"id"` // 要清除的 ScopePermission ID
}

// ScopesInvalidateCacheEvent Scopes 缓存清除事件
type ScopesInvalidateCacheEvent struct {
	events.AccessTopic
	ID string `json:"id"` // 要清除的 Scope ID (通常是 scope 字符串)
}
