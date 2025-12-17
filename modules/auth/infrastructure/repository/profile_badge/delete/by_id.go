package delete

import (
	"context"
	profileBadgeDomainErrors "nfxid/modules/auth/domain/profile_badge/errors"
	"nfxid/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByID 根据 ID 删除 ProfileBadge，实现 profileBadge.Delete 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	result := h.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.ProfileBadge{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return profileBadgeDomainErrors.ErrProfileBadgeNotFound
	}
	return nil
}
