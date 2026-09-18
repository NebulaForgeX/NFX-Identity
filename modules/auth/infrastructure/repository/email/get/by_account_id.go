package get

import (
	"context"

	emailDomain "nfxidentity/modules/auth/domain/email"
	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/email/mapper"

	"github.com/google/uuid"
)

func (h *Handler) ByAccountID(ctx context.Context, accountID uuid.UUID) ([]*emailDomain.Email, error) {
	var ms []rdbmodels.Email
	if err := h.db.WithContext(ctx).
		Where(rdbmodels.EmailCols.AccountID+" = ?", accountID.String()).
		Order(rdbmodels.EmailCols.IsPrimary + " DESC").
		Order(rdbmodels.EmailCols.CreatedAt + " ASC").
		Find(&ms).Error; err != nil {
		return nil, err
	}
	return mapper.EmailModelsToDomains(ms), nil
}
