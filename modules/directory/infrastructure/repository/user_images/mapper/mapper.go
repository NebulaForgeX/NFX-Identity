package mapper

import (
	"nfxid/modules/directory/domain/user_images"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/pkgs/utils/timex"
)

// UserImageDomainToModel 将 Domain UserImage 转换为 Model UserImage
func UserImageDomainToModel(ui *user_images.UserImage) *models.UserImage {
	if ui == nil {
		return nil
	}

	return &models.UserImage{
		ID:           ui.ID(),
		UserID:       ui.UserID(),
		ImageID:      ui.ImageID(),
		DisplayOrder: ui.DisplayOrder(),
		CreatedAt:    ui.CreatedAt(),
		UpdatedAt:    ui.UpdatedAt(),
		DeletedAt:    timex.TimeToGormDeletedAt(ui.DeletedAt()),
	}
}

// UserImageModelToDomain 将 Model UserImage 转换为 Domain UserImage
func UserImageModelToDomain(m *models.UserImage) *user_images.UserImage {
	if m == nil {
		return nil
	}

	state := user_images.UserImageState{
		ID:           m.ID,
		UserID:       m.UserID,
		ImageID:      m.ImageID,
		DisplayOrder: m.DisplayOrder,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
		DeletedAt:    timex.GormDeletedAtToTime(m.DeletedAt),
	}

	return user_images.NewUserImageFromState(state)
}

// UserImageModelToUpdates 将 Model UserImage 转换为更新字段映射
func UserImageModelToUpdates(m *models.UserImage) map[string]any {
	return map[string]any{
		models.UserImageCols.UserID:       m.UserID,
		models.UserImageCols.ImageID:      m.ImageID,
		models.UserImageCols.DisplayOrder: m.DisplayOrder,
		models.UserImageCols.UpdatedAt:    m.UpdatedAt,
		models.UserImageCols.DeletedAt:    m.DeletedAt,
	}
}
