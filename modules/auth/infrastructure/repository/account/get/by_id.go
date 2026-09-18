package get

import (
	"context"
	"errors"

	authErr "nfxidentity/errors/src/auth"
	accountDomain "nfxidentity/modules/auth/domain/account"
	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/account/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*accountDomain.Account, error) {
	var m rdbmodels.Account
	if err := h.db.WithContext(ctx).Where(rdbmodels.AccountCols.ID+" = ?", id.String()).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, authErr.ErrAccountNotFound
		}
		return nil, err
	}
	return mapper.AccountModelToDomain(&m), nil
}
