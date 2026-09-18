package get

import (
	"context"
	"errors"

	"nfxidentity/errors/src/auth"
	"nfxidentity/modules/auth/domain/phone"
	"nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/phone/mapper"

	"gorm.io/gorm"
)

func (h *Handler) ByNumber(ctx context.Context, number string) (*phone.Phone, error) {
	var m models.Phone
	if err := h.db.WithContext(ctx).Where("phone = ? AND deleted_at IS NULL", number).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, auth.ErrPhoneBindingNotFound
		}
		return nil, err
	}
	return mapper.ToDomain(&m), nil
}
