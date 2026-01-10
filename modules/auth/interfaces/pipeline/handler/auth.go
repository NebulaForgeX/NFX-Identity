package handler

import (
	"context"

	"nfxid/events/auth"
	"nfxid/pkgs/logx"

	"github.com/ThreeDotsLabs/watermill/message"
)

type AuthHandler struct {
	// 可以注入 application services 或其他依赖
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

// OnAccountLockoutsInvalidateCache 监听 AccountLockouts 缓存清除事件
func (h *AuthHandler) OnAccountLockoutsInvalidateCache(ctx context.Context, evt auth.AccountLockoutsInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Auth Pipeline] 收到 AccountLockouts 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	// 缓存清除事件处理完成
	return nil
}

// OnLoginAttemptsInvalidateCache 监听 LoginAttempts 缓存清除事件
func (h *AuthHandler) OnLoginAttemptsInvalidateCache(ctx context.Context, evt auth.LoginAttemptsInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Auth Pipeline] 收到 LoginAttempts 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	// 缓存清除事件处理完成
	return nil
}

// OnMFAFactorsInvalidateCache 监听 MFAFactors 缓存清除事件
func (h *AuthHandler) OnMFAFactorsInvalidateCache(ctx context.Context, evt auth.MFAFactorsInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Auth Pipeline] 收到 MFAFactors 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	// 缓存清除事件处理完成
	return nil
}

// OnPasswordHistoryInvalidateCache 监听 PasswordHistory 缓存清除事件
func (h *AuthHandler) OnPasswordHistoryInvalidateCache(ctx context.Context, evt auth.PasswordHistoryInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Auth Pipeline] 收到 PasswordHistory 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	// 缓存清除事件处理完成
	return nil
}

// OnPasswordResetsInvalidateCache 监听 PasswordResets 缓存清除事件
func (h *AuthHandler) OnPasswordResetsInvalidateCache(ctx context.Context, evt auth.PasswordResetsInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Auth Pipeline] 收到 PasswordResets 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	// 缓存清除事件处理完成
	return nil
}

// OnRefreshTokensInvalidateCache 监听 RefreshTokens 缓存清除事件
func (h *AuthHandler) OnRefreshTokensInvalidateCache(ctx context.Context, evt auth.RefreshTokensInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Auth Pipeline] 收到 RefreshTokens 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	// 缓存清除事件处理完成
	return nil
}

// OnSessionsInvalidateCache 监听 Sessions 缓存清除事件
func (h *AuthHandler) OnSessionsInvalidateCache(ctx context.Context, evt auth.SessionsInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Auth Pipeline] 收到 Sessions 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	// 缓存清除事件处理完成
	return nil
}

// OnTrustedDevicesInvalidateCache 监听 TrustedDevices 缓存清除事件
func (h *AuthHandler) OnTrustedDevicesInvalidateCache(ctx context.Context, evt auth.TrustedDevicesInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Auth Pipeline] 收到 TrustedDevices 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	// 缓存清除事件处理完成
	return nil
}

// OnUserCredentialsInvalidateCache 监听 UserCredentials 缓存清除事件
func (h *AuthHandler) OnUserCredentialsInvalidateCache(ctx context.Context, evt auth.UserCredentialsInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Auth Pipeline] 收到 UserCredentials 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	// 缓存清除事件处理完成
	return nil
}
