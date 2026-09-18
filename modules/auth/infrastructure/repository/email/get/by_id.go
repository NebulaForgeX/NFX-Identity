package get

import (
	"context"
	"errors"

	authErr "nfxidentity/errors/src/auth"
	emailDomain "nfxidentity/modules/auth/domain/email"
	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/email/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*emailDomain.Email, error) {
	var m rdbmodels.Email
	if err := h.db.WithContext(ctx).Where(rdbmodels.EmailCols.ID+" = ?", id.String()).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, authErr.ErrEmailBindingNotFound
		}
		return nil, err
	}
	return mapper.EmailModelToDomain(&m), nil
}
