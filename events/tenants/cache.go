package tenants

import (
	"nfxid/events"
)

// DomainVerificationsInvalidateCacheEvent DomainVerifications 缓存清除事件
// EventType 会自动从类型名生成，TopicKey 通过嵌入 events.TenantsTopic 自动提供
// Cache key 格式: {prefix[:namespace]}:entity:{id}
type DomainVerificationsInvalidateCacheEvent struct {
	events.TenantsTopic
	ID        string `json:"id"`         // 要清除的 Domain Verification ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "domain_verification"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// GroupsInvalidateCacheEvent Groups 缓存清除事件
type GroupsInvalidateCacheEvent struct {
	events.TenantsTopic
	ID        string `json:"id"`         // 要清除的 Group ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "group"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// InvitationsInvalidateCacheEvent Invitations 缓存清除事件
type InvitationsInvalidateCacheEvent struct {
	events.TenantsTopic
	ID        string `json:"id"`         // 要清除的 Invitation ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "invitation"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// MemberAppRolesInvalidateCacheEvent MemberAppRoles 缓存清除事件
type MemberAppRolesInvalidateCacheEvent struct {
	events.TenantsTopic
	ID        string `json:"id"`         // 要清除的 Member App Role ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "member_app_role"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// MemberGroupsInvalidateCacheEvent MemberGroups 缓存清除事件
type MemberGroupsInvalidateCacheEvent struct {
	events.TenantsTopic
	ID        string `json:"id"`         // 要清除的 Member Group ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "member_group"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// MemberRolesInvalidateCacheEvent MemberRoles 缓存清除事件
type MemberRolesInvalidateCacheEvent struct {
	events.TenantsTopic
	ID        string `json:"id"`         // 要清除的 Member Role ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "member_role"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// MembersInvalidateCacheEvent Members 缓存清除事件
type MembersInvalidateCacheEvent struct {
	events.TenantsTopic
	ID        string `json:"id"`         // 要清除的 Member ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "member"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// TenantAppsInvalidateCacheEvent TenantApps 缓存清除事件
type TenantAppsInvalidateCacheEvent struct {
	events.TenantsTopic
	ID        string `json:"id"`         // 要清除的 Tenant App ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "tenant_app"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// TenantSettingsInvalidateCacheEvent TenantSettings 缓存清除事件
type TenantSettingsInvalidateCacheEvent struct {
	events.TenantsTopic
	ID        string `json:"id"`         // 要清除的 Tenant Setting ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "tenant_setting"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// TenantsInvalidateCacheEvent Tenants 缓存清除事件
type TenantsInvalidateCacheEvent struct {
	events.TenantsTopic
	ID        string `json:"id"`         // 要清除的 Tenant ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "tenant"
	Namespace string `json:"namespace"` // Cache namespace，可选
}
