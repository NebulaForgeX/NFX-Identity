package tenants

import (
	"nfxid/events"
)

// DomainVerificationsInvalidateCacheEvent DomainVerifications 缓存清除事件
// EventType 会自动从类型名生成，TopicKey 通过嵌入 events.TenantsTopic 自动提供
type DomainVerificationsInvalidateCacheEvent struct {
	events.TenantsTopic
	ID string `json:"id"` // 要清除的 Domain Verification ID
}

// GroupsInvalidateCacheEvent Groups 缓存清除事件
type GroupsInvalidateCacheEvent struct {
	events.TenantsTopic
	ID string `json:"id"` // 要清除的 Group ID
}

// InvitationsInvalidateCacheEvent Invitations 缓存清除事件
type InvitationsInvalidateCacheEvent struct {
	events.TenantsTopic
	ID string `json:"id"` // 要清除的 Invitation ID
}

// MemberAppRolesInvalidateCacheEvent MemberAppRoles 缓存清除事件
type MemberAppRolesInvalidateCacheEvent struct {
	events.TenantsTopic
	ID string `json:"id"` // 要清除的 Member App Role ID
}

// MemberGroupsInvalidateCacheEvent MemberGroups 缓存清除事件
type MemberGroupsInvalidateCacheEvent struct {
	events.TenantsTopic
	ID string `json:"id"` // 要清除的 Member Group ID
}

// MemberRolesInvalidateCacheEvent MemberRoles 缓存清除事件
type MemberRolesInvalidateCacheEvent struct {
	events.TenantsTopic
	ID string `json:"id"` // 要清除的 Member Role ID
}

// MembersInvalidateCacheEvent Members 缓存清除事件
type MembersInvalidateCacheEvent struct {
	events.TenantsTopic
	ID string `json:"id"` // 要清除的 Member ID
}

// TenantAppsInvalidateCacheEvent TenantApps 缓存清除事件
type TenantAppsInvalidateCacheEvent struct {
	events.TenantsTopic
	ID string `json:"id"` // 要清除的 Tenant App ID
}

// TenantSettingsInvalidateCacheEvent TenantSettings 缓存清除事件
type TenantSettingsInvalidateCacheEvent struct {
	events.TenantsTopic
	ID string `json:"id"` // 要清除的 Tenant Setting ID
}

// TenantsInvalidateCacheEvent Tenants 缓存清除事件
type TenantsInvalidateCacheEvent struct {
	events.TenantsTopic
	ID string `json:"id"` // 要清除的 Tenant ID
}
