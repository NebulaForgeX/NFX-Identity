package handler

import (
	"context"

	"nfxid/events"
	badgeApp "nfxid/modules/auth/application/badge"
	profileApp "nfxid/modules/auth/application/profile"
	profileBadgeApp "nfxid/modules/auth/application/profile_badge"
	educationApp "nfxid/modules/auth/application/profile_education"
	occupationApp "nfxid/modules/auth/application/profile_occupation"
	roleApp "nfxid/modules/auth/application/role"
	userApp "nfxid/modules/auth/application/user"
	"nfxid/pkgs/logx"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
)

type AuthHandler struct {
	userAppSvc         *userApp.Service
	profileAppSvc      *profileApp.Service
	roleAppSvc         *roleApp.Service
	badgeAppSvc        *badgeApp.Service
	educationAppSvc    *educationApp.Service
	occupationAppSvc   *occupationApp.Service
	profileBadgeAppSvc *profileBadgeApp.Service
}

func NewAuthHandler(
	userAppSvc *userApp.Service,
	profileAppSvc *profileApp.Service,
	roleAppSvc *roleApp.Service,
	badgeAppSvc *badgeApp.Service,
	educationAppSvc *educationApp.Service,
	occupationAppSvc *occupationApp.Service,
	profileBadgeAppSvc *profileBadgeApp.Service,
) *AuthHandler {
	return &AuthHandler{
		userAppSvc:         userAppSvc,
		profileAppSvc:      profileAppSvc,
		roleAppSvc:         roleAppSvc,
		badgeAppSvc:        badgeAppSvc,
		educationAppSvc:    educationAppSvc,
		occupationAppSvc:   occupationAppSvc,
		profileBadgeAppSvc: profileBadgeAppSvc,
	}
}

// OnAuthToAuth_Success 监听 Auth 内部成功消息
func (h *AuthHandler) OnAuthToAuth_Success(ctx context.Context, evt events.AuthToAuth_SuccessEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Auth Worker] 已收到成功消息: operation=%s, entity_id=%s, user_id=%s, details=%+v",
		evt.Operation, evt.EntityID, evt.UserID, evt.Details)

	// 通用成功事件处理器，用于日志记录和后续扩展
	// 如需添加业务逻辑（如更新缓存、发送通知等），可在此处扩展
	return nil
}

// OnAuthToAuth_Test 监听 Auth 内部测试消息
func (h *AuthHandler) OnAuthToAuth_Test(ctx context.Context, evt events.AuthToAuth_TestEvent, msg *message.Message) error {
	logx.S().Infof("📨 [Auth Worker] 收到 Auth 测试消息: %s", evt.Message)
	return nil
}

// OnAuthToAuth_UserInvalidateCache 监听用户缓存清除事件（Auth 内部）
func (h *AuthHandler) OnAuthToAuth_UserInvalidateCache(ctx context.Context, evt events.AuthToAuth_UserInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("🗑️ [Auth Worker] 清除用户缓存: user_id=%s, operation=%s", evt.UserID, evt.Operation)
	// 注意：User service 当前没有缓存，正常查询都没有缓存
	// 如果需要添加缓存，可以在这里调用缓存清理逻辑
	return nil
}

// OnAuthToAuth_ProfileInvalidateCache 监听资料缓存清除事件（Auth 内部）
func (h *AuthHandler) OnAuthToAuth_ProfileInvalidateCache(ctx context.Context, evt events.AuthToAuth_ProfileInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("🗑️ [Auth Worker] 清除资料缓存: profile_id=%s, user_id=%s, operation=%s", evt.ProfileID, evt.UserID, evt.Operation)
	// 注意：Profile service 当前没有缓存，正常查询都没有缓存
	// 如果需要添加缓存，可以在这里调用缓存清理逻辑
	return nil
}

// OnAuthToAuth_RoleInvalidateCache 监听角色缓存清除事件（Auth 内部）
func (h *AuthHandler) OnAuthToAuth_RoleInvalidateCache(ctx context.Context, evt events.AuthToAuth_RoleInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("🗑️ [Auth Worker] 清除角色缓存: role_id=%s, operation=%s", evt.RoleID, evt.Operation)
	// 注意：Role service 当前没有缓存，正常查询都没有缓存
	// 如果需要添加缓存，可以在这里调用缓存清理逻辑
	return nil
}

// OnAuthToAuth_BadgeInvalidateCache 监听徽章缓存清除事件（Auth 内部）
func (h *AuthHandler) OnAuthToAuth_BadgeInvalidateCache(ctx context.Context, evt events.AuthToAuth_BadgeInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("🗑️ [Auth Worker] 清除徽章缓存: badge_id=%s, operation=%s", evt.BadgeID, evt.Operation)

	badgeID, err := uuid.Parse(evt.BadgeID)
	if err != nil {
		logx.S().Warnf("无效的 badge_id: %s, error: %v", evt.BadgeID, err)
		return nil
	}

	// Badge service 有缓存，调用缓存清理
	if err := h.badgeAppSvc.InvalidateBadgeCache(ctx, badgeID); err != nil {
		logx.S().Errorf("清除徽章缓存失败: badge_id=%s, error: %v", evt.BadgeID, err)
		return err
	}

	return nil
}

// OnAuthToAuth_EducationInvalidateCache 监听教育经历缓存清除事件（Auth 内部）
func (h *AuthHandler) OnAuthToAuth_EducationInvalidateCache(ctx context.Context, evt events.AuthToAuth_EducationInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("🗑️ [Auth Worker] 清除教育经历缓存: education_id=%s, profile_id=%s, operation=%s", evt.EducationID, evt.ProfileID, evt.Operation)

	educationID, err := uuid.Parse(evt.EducationID)
	if err != nil {
		logx.S().Warnf("无效的 education_id: %s, error: %v", evt.EducationID, err)
		return nil
	}

	// Education service 有缓存，调用缓存清理
	if err := h.educationAppSvc.InvalidateEducationCache(ctx, educationID); err != nil {
		logx.S().Errorf("清除教育经历缓存失败: education_id=%s, error: %v", evt.EducationID, err)
		return err
	}

	return nil
}

// OnAuthToAuth_OccupationInvalidateCache 监听职业信息缓存清除事件（Auth 内部）
func (h *AuthHandler) OnAuthToAuth_OccupationInvalidateCache(ctx context.Context, evt events.AuthToAuth_OccupationInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("🗑️ [Auth Worker] 清除职业信息缓存: occupation_id=%s, profile_id=%s, operation=%s", evt.OccupationID, evt.ProfileID, evt.Operation)
	// 注意：Occupation service 当前没有缓存，正常查询都没有缓存
	// 如果需要添加缓存，可以在这里调用缓存清理逻辑
	return nil
}

// OnAuthToAuth_ProfileBadgeInvalidateCache 监听用户徽章关联缓存清除事件（Auth 内部）
func (h *AuthHandler) OnAuthToAuth_ProfileBadgeInvalidateCache(ctx context.Context, evt events.AuthToAuth_ProfileBadgeInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("🗑️ [Auth Worker] 清除用户徽章关联缓存: profile_badge_id=%s, profile_id=%s, badge_id=%s, operation=%s", evt.ProfileBadgeID, evt.ProfileID, evt.BadgeID, evt.Operation)
	// 注意：ProfileBadge service 当前没有缓存，正常查询都没有缓存
	// 如果需要添加缓存，可以在这里调用缓存清理逻辑
	return nil
}

// OnAuthToAuth_UserCreated 监听用户创建事件（Auth 内部，用于通知其他服务创建关联数据）
func (h *AuthHandler) OnAuthToAuth_UserCreated(ctx context.Context, evt events.AuthToAuth_UserCreatedEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Auth Worker] 收到用户创建事件: user_id=%s, username=%s, email=%s, status=%s", evt.UserID, evt.Username, evt.Email, evt.Status)

	// 事件已发布，profile 等服务可通过监听此事件来处理关联数据的创建
	// 保持解耦设计：profile 的创建由业务逻辑决定（例如注册时通过 HTTP handler 或 gRPC 调用同时创建）
	return nil
}

// OnAuthToAuth_UserUpdated 监听用户更新事件（Auth 内部，用于通知其他服务）
func (h *AuthHandler) OnAuthToAuth_UserUpdated(ctx context.Context, evt events.AuthToAuth_UserUpdatedEvent, msg *message.Message) error {
	logx.S().Infof("📝 [Auth Worker] 收到用户更新事件: user_id=%s, username=%s, email=%s", evt.UserID, evt.Username, evt.Email)

	// 事件已发布，其他服务可通过监听此事件来处理关联数据的更新
	// 保持解耦设计：关联数据的同步由业务逻辑决定（可通过 gRPC 调用或事件驱动）

	return nil
}

// OnAuthToAuth_UserDeleted 监听用户删除事件（Auth 内部，用于通知其他服务删除关联数据）
func (h *AuthHandler) OnAuthToAuth_UserDeleted(ctx context.Context, evt events.AuthToAuth_UserDeletedEvent, msg *message.Message) error {
	logx.S().Infof("🗑️ [Auth Worker] 收到用户删除事件: user_id=%s, username=%s, email=%s", evt.UserID, evt.Username, evt.Email)

	userID, err := uuid.Parse(evt.UserID)
	if err != nil {
		logx.S().Warnf("无效的 user_id: %s, error: %v", evt.UserID, err)
		return nil
	}

	// 删除关联的 profile（如果存在）
	if err := h.profileAppSvc.DeleteByUserID(ctx, userID); err != nil {
		logx.S().Errorf("删除用户资料失败: user_id=%s, error: %v", evt.UserID, err)
		// 不返回错误，避免影响其他服务的处理
	}

	return nil
}
