package get

import (
	"context"
	"errors"

	"nfxidentity/errors/src/auth"
	"nfxidentity/modules/auth/domain/forgerprofile"
	"nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/forgerprofile/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (h *Handler) ByAccountAndID(ctx context.Context, accountID, id uuid.UUID) (*forgerprofile.Profile, error) {
	var m models.Forgerprofile
	if err := h.db.WithContext(ctx).Where("id = ? AND account_id = ? AND deleted_at IS NULL", id, accountID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, auth.ErrForgerProfileNotFound
		}
		return nil, err
	}
	return mapper.ToDomain(&m), nil
}
