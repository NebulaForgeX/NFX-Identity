package get

import (
	"context"
	"errors"

	"nfxidentity/errors/src/auth"
	"nfxidentity/modules/auth/domain/email"
	"nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/email/mapper"

	"gorm.io/gorm"
)

func (h *Handler) ByAddress(ctx context.Context, address string) (*email.Email, error) {
	var m models.Email
	if err := h.db.WithContext(ctx).Where("email = ? AND deleted_at IS NULL", address).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, auth.ErrEmailBindingNotFound
		}
		return nil, err
	}
	return mapper.ToDomain(&m), nil
}
