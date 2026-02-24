package get

import (
	"context"
	"errors"

	accessErr "nfxid/errors/src/access"
	"nfxid/modules/access/domain/super_admins"
	"nfxid/modules/access/infrastructure/rdb/models"
	"nfxid/modules/access/infrastructure/repository/super_admins/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (h *Handler) ByUserID(ctx context.Context, userID uuid.UUID) (*super_admins.SuperAdmin, error) {
	var m models.SuperAdmin
	if err := h.db.WithContext(ctx).Where("user_id = ?", userID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, accessErr.ErrSuperAdminNotFound
		}
		return nil, err
	}
	return mapper.SuperAdminModelToDomain(&m), nil
}
