package get

import (
	"context"
	"errors"
	"nfxid/modules/directory/domain/badges"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/modules/directory/infrastructure/repository/badges/mapper"

	"gorm.io/gorm"
)

// ByName 根据 Name 获取 Badge，实现 badges.Get 接口
func (h *Handler) ByName(ctx context.Context, name string) (*badges.Badge, error) {
	var m models.Badge
	if err := h.db.WithContext(ctx).Where("name = ?", name).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, badges.ErrBadgeNotFound
		}
		return nil, err
	}
	return mapper.BadgeModelToDomain(&m), nil
}
