package handler

import (
	"context"

	"nfxid/events/access"
	"nfxid/pkgs/logx"

	"github.com/ThreeDotsLabs/watermill/message"
)

type AccessHandler struct {
	// 可以注入 application services 或其他依赖
}

func NewAccessHandler() *AccessHandler {
	return &AccessHandler{}
}

// OnGrantsInvalidateCache 监听 Grants 缓存清除事件
func (h *AccessHandler) OnGrantsInvalidateCache(ctx context.Context, evt access.GrantsInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Access Pipeline] 收到 Grants 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	// 缓存清除事件处理完成
	return nil
}

// OnPermissionsInvalidateCache 监听 Permissions 缓存清除事件
func (h *AccessHandler) OnPermissionsInvalidateCache(ctx context.Context, evt access.PermissionsInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Access Pipeline] 收到 Permissions 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	// 缓存清除事件处理完成
	return nil
}

// OnRolesInvalidateCache 监听 Roles 缓存清除事件
func (h *AccessHandler) OnRolesInvalidateCache(ctx context.Context, evt access.RolesInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Access Pipeline] 收到 Roles 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	// 缓存清除事件处理完成
	return nil
}

// OnScopesInvalidateCache 监听 Scopes 缓存清除事件
func (h *AccessHandler) OnScopesInvalidateCache(ctx context.Context, evt access.ScopesInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Access Pipeline] 收到 Scopes 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	// 缓存清除事件处理完成
	return nil
}

// OnRolePermissionsInvalidateCache 监听 RolePermissions 缓存清除事件
func (h *AccessHandler) OnRolePermissionsInvalidateCache(ctx context.Context, evt access.RolePermissionsInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Access Pipeline] 收到 RolePermissions 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	// 缓存清除事件处理完成
	return nil
}

// OnScopePermissionsInvalidateCache 监听 ScopePermissions 缓存清除事件
func (h *AccessHandler) OnScopePermissionsInvalidateCache(ctx context.Context, evt access.ScopePermissionsInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Access Pipeline] 收到 ScopePermissions 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	// 缓存清除事件处理完成
	return nil
}
