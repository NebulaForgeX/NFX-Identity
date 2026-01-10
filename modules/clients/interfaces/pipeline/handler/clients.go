package handler

import (
	"context"

	"nfxid/events/clients"
	"nfxid/pkgs/logx"

	"github.com/ThreeDotsLabs/watermill/message"
)

type ClientsHandler struct {
}

func NewClientsHandler() *ClientsHandler {
	return &ClientsHandler{}
}

// OnAPIKeysInvalidateCache 监听 APIKeys 缓存清除事件
func (h *ClientsHandler) OnAPIKeysInvalidateCache(ctx context.Context, evt clients.APIKeysInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Clients Pipeline] 收到 APIKeys 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	return nil
}

// OnAppsInvalidateCache 监听 Apps 缓存清除事件
func (h *ClientsHandler) OnAppsInvalidateCache(ctx context.Context, evt clients.AppsInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Clients Pipeline] 收到 Apps 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	return nil
}

// OnClientCredentialsInvalidateCache 监听 ClientCredentials 缓存清除事件
func (h *ClientsHandler) OnClientCredentialsInvalidateCache(ctx context.Context, evt clients.ClientCredentialsInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Clients Pipeline] 收到 ClientCredentials 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	return nil
}

// OnClientScopesInvalidateCache 监听 ClientScopes 缓存清除事件
func (h *ClientsHandler) OnClientScopesInvalidateCache(ctx context.Context, evt clients.ClientScopesInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Clients Pipeline] 收到 ClientScopes 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	return nil
}

// OnIPAllowlistInvalidateCache 监听 IPAllowlist 缓存清除事件
func (h *ClientsHandler) OnIPAllowlistInvalidateCache(ctx context.Context, evt clients.IPAllowlistInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Clients Pipeline] 收到 IPAllowlist 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	return nil
}

// OnRateLimitsInvalidateCache 监听 RateLimits 缓存清除事件
func (h *ClientsHandler) OnRateLimitsInvalidateCache(ctx context.Context, evt clients.RateLimitsInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Clients Pipeline] 收到 RateLimits 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	return nil
}
