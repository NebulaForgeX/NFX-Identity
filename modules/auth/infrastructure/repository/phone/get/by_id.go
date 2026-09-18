package get

import (
	"context"
	"errors"

	"nfxidentity/errors/src/auth"
	"nfxidentity/modules/auth/domain/phone"
	"nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/phone/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*phone.Phone, error) {
	var m models.Phone
	if err := h.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, auth.ErrPhoneBindingNotFound
		}
		return nil, err
	}
	return mapper.ToDomain(&m), nil
}
