package get

import (
	"context"
	"errors"

	authErr "nfxidentity/errors/src/auth"
	phoneDomain "nfxidentity/modules/auth/domain/phone"
	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/phone/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*phoneDomain.Phone, error) {
	var m rdbmodels.Phone
	if err := h.db.WithContext(ctx).Where(rdbmodels.PhoneCols.ID+" = ?", id.String()).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, authErr.ErrPhoneBindingNotFound
		}
		return nil, err
	}
	return mapper.PhoneModelToDomain(&m), nil
}
