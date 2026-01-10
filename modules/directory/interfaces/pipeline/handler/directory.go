package handler

import (
	"context"

	"nfxid/events/directory"
	"nfxid/pkgs/logx"

	"github.com/ThreeDotsLabs/watermill/message"
)

type DirectoryHandler struct {
	// 可以注入 application services 或其他依赖
}

func NewDirectoryHandler() *DirectoryHandler {
	return &DirectoryHandler{}
}

// OnUsersInvalidateCache 监听 Users 缓存清除事件
func (h *DirectoryHandler) OnUsersInvalidateCache(ctx context.Context, evt directory.UsersInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Directory Pipeline] 收到 Users 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	return nil
}

// OnBadgesInvalidateCache 监听 Badges 缓存清除事件
func (h *DirectoryHandler) OnBadgesInvalidateCache(ctx context.Context, evt directory.BadgesInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Directory Pipeline] 收到 Badges 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	return nil
}

// OnUserBadgesInvalidateCache 监听 UserBadges 缓存清除事件
func (h *DirectoryHandler) OnUserBadgesInvalidateCache(ctx context.Context, evt directory.UserBadgesInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Directory Pipeline] 收到 UserBadges 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	return nil
}

// OnUserEducationsInvalidateCache 监听 UserEducations 缓存清除事件
func (h *DirectoryHandler) OnUserEducationsInvalidateCache(ctx context.Context, evt directory.UserEducationsInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Directory Pipeline] 收到 UserEducations 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	return nil
}

// OnUserEmailsInvalidateCache 监听 UserEmails 缓存清除事件
func (h *DirectoryHandler) OnUserEmailsInvalidateCache(ctx context.Context, evt directory.UserEmailsInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Directory Pipeline] 收到 UserEmails 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	return nil
}

// OnUserOccupationsInvalidateCache 监听 UserOccupations 缓存清除事件
func (h *DirectoryHandler) OnUserOccupationsInvalidateCache(ctx context.Context, evt directory.UserOccupationsInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Directory Pipeline] 收到 UserOccupations 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	return nil
}

// OnUserPhonesInvalidateCache 监听 UserPhones 缓存清除事件
func (h *DirectoryHandler) OnUserPhonesInvalidateCache(ctx context.Context, evt directory.UserPhonesInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Directory Pipeline] 收到 UserPhones 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	return nil
}

// OnUserPreferencesInvalidateCache 监听 UserPreferences 缓存清除事件
func (h *DirectoryHandler) OnUserPreferencesInvalidateCache(ctx context.Context, evt directory.UserPreferencesInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Directory Pipeline] 收到 UserPreferences 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	return nil
}

// OnUserProfilesInvalidateCache 监听 UserProfiles 缓存清除事件
func (h *DirectoryHandler) OnUserProfilesInvalidateCache(ctx context.Context, evt directory.UserProfilesInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Directory Pipeline] 收到 UserProfiles 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	return nil
}
