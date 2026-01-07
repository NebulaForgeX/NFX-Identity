package get

import (
	"context"
	"errors"
	profileBadge "nfxid/modules/auth/domain/profile_badge"
	profileBadgeDomainErrors "nfxid/modules/auth/domain/profile_badge/errors"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 ProfileBadge，实现 profileBadge.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*profileBadge.ProfileBadge, error) {
	var m models.ProfileBadge
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, profileBadgeDomainErrors.ErrProfileBadgeNotFound
		}
		return nil, err
	}
	return mapper.ProfileBadgeModelToDomain(&m), nil
}
