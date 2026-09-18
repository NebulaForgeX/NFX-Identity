package get

import (
	"context"
	"errors"

	authErr "nfxidentity/errors/src/auth"
	phoneDomain "nfxidentity/modules/auth/domain/phone"
	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/phone/mapper"

	"gorm.io/gorm"
)

func (h *Handler) ByPhone(ctx context.Context, canonicalPhone string) (*phoneDomain.Phone, error) {
	var m rdbmodels.Phone
	if err := h.db.WithContext(ctx).Where(rdbmodels.PhoneCols.Phone+" = ?", canonicalPhone).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, authErr.ErrPhoneBindingNotFound
		}
		return nil, err
	}
	return mapper.PhoneModelToDomain(&m), nil
}
