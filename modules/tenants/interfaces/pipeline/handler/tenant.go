package handler

import (
	"context"

	"nfxid/events/tenants"
	"nfxid/pkgs/logx"

	"github.com/ThreeDotsLabs/watermill/message"
)

type TenantsHandler struct {
	// 可以注入 application services 或其他依赖
}

func NewTenantsHandler() *TenantsHandler {
	return &TenantsHandler{}
}

// OnTenantsInvalidateCache 监听 Tenants 缓存清除事件
func (h *TenantsHandler) OnTenantsInvalidateCache(ctx context.Context, evt tenants.TenantsInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Tenants Pipeline] 收到 Tenants 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	return nil
}

// OnGroupsInvalidateCache 监听 Groups 缓存清除事件
func (h *TenantsHandler) OnGroupsInvalidateCache(ctx context.Context, evt tenants.GroupsInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Tenants Pipeline] 收到 Groups 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	return nil
}

// OnMembersInvalidateCache 监听 Members 缓存清除事件
func (h *TenantsHandler) OnMembersInvalidateCache(ctx context.Context, evt tenants.MembersInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Tenants Pipeline] 收到 Members 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	return nil
}

// OnInvitationsInvalidateCache 监听 Invitations 缓存清除事件
func (h *TenantsHandler) OnInvitationsInvalidateCache(ctx context.Context, evt tenants.InvitationsInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Tenants Pipeline] 收到 Invitations 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	return nil
}

// OnTenantAppsInvalidateCache 监听 TenantApps 缓存清除事件
func (h *TenantsHandler) OnTenantAppsInvalidateCache(ctx context.Context, evt tenants.TenantAppsInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Tenants Pipeline] 收到 TenantApps 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	return nil
}

// OnTenantSettingsInvalidateCache 监听 TenantSettings 缓存清除事件
func (h *TenantsHandler) OnTenantSettingsInvalidateCache(ctx context.Context, evt tenants.TenantSettingsInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Tenants Pipeline] 收到 TenantSettings 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	return nil
}

// OnDomainVerificationsInvalidateCache 监听 DomainVerifications 缓存清除事件
func (h *TenantsHandler) OnDomainVerificationsInvalidateCache(ctx context.Context, evt tenants.DomainVerificationsInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Tenants Pipeline] 收到 DomainVerifications 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	return nil
}

// OnMemberRolesInvalidateCache 监听 MemberRoles 缓存清除事件
func (h *TenantsHandler) OnMemberRolesInvalidateCache(ctx context.Context, evt tenants.MemberRolesInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Tenants Pipeline] 收到 MemberRoles 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	return nil
}

// OnMemberGroupsInvalidateCache 监听 MemberGroups 缓存清除事件
func (h *TenantsHandler) OnMemberGroupsInvalidateCache(ctx context.Context, evt tenants.MemberGroupsInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Tenants Pipeline] 收到 MemberGroups 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	return nil
}

// OnMemberAppRolesInvalidateCache 监听 MemberAppRoles 缓存清除事件
func (h *TenantsHandler) OnMemberAppRolesInvalidateCache(ctx context.Context, evt tenants.MemberAppRolesInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Tenants Pipeline] 收到 MemberAppRoles 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	return nil
}
