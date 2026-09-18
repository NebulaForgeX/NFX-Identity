package get

import (
	"context"
	"errors"

	"nfxidentity/errors/src/auth"
	"nfxidentity/modules/auth/domain/account"
	"nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/account/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*account.Account, error) {
	var m models.Account
	if err := h.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, auth.ErrAccountNotFound
		}
		return nil, err
	}
	return mapper.ToDomain(&m), nil
}
