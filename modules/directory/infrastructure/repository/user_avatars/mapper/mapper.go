package mapper

import (
	"nfxid/modules/directory/domain/user_avatars"
	"nfxid/modules/directory/infrastructure/rdb/models"
)

// UserAvatarDomainToModel 将 Domain UserAvatar 转换为 Model UserAvatar
func UserAvatarDomainToModel(ua *user_avatars.UserAvatar) *models.UserAvatar {
	if ua == nil {
		return nil
	}

	return &models.UserAvatar{
		UserID:    ua.UserID(),
		ImageID:   ua.ImageID(),
		CreatedAt: ua.CreatedAt(),
		UpdatedAt: ua.UpdatedAt(),
	}
}

// UserAvatarModelToDomain 将 Model UserAvatar 转换为 Domain UserAvatar
func UserAvatarModelToDomain(m *models.UserAvatar) *user_avatars.UserAvatar {
	if m == nil {
		return nil
	}

	state := user_avatars.UserAvatarState{
		UserID:    m.UserID,
		ImageID:   m.ImageID,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}

	return user_avatars.NewUserAvatarFromState(state)
}

// UserAvatarModelToUpdates 将 Model UserAvatar 转换为更新字段映射
func UserAvatarModelToUpdates(m *models.UserAvatar) map[string]any {
	return map[string]any{
		models.UserAvatarCols.ImageID:   m.ImageID,
		models.UserAvatarCols.UpdatedAt: m.UpdatedAt,
	}
}
