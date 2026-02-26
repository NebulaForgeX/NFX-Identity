package create

import (
	"context"
	"nfxidentity/modules/directory/domain/user_images"
	"nfxidentity/modules/directory/infrastructure/repository/user_images/mapper"
)

// New 创建新的 UserImage，实现 user_images.Create 接口
func (h *Handler) New(ctx context.Context, ui *user_images.UserImage) error {
	m := mapper.UserImageDomainToModel(ui)
	return h.db.WithContext(ctx).Create(&m).Error
}
