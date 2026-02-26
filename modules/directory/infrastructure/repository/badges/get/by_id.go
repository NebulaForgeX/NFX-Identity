package get

import (
	"context"
	"errors"

	dirErr "nfxidentity/errors/src/directory"
	"nfxidentity/modules/directory/domain/badges"
	"nfxidentity/modules/directory/infrastructure/rdb/models"
	"nfxidentity/modules/directory/infrastructure/repository/badges/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 Badge，实现 badges.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*badges.Badge, error) {
	var m models.Badge
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, dirErr.ErrBadgeNotFound
		}
		return nil, err
	}
	return mapper.BadgeModelToDomain(&m), nil
}
