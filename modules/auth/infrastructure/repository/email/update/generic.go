package update

import (
	"context"

	authErr "nfxidentity/errors/src/auth"
	emailDomain "nfxidentity/modules/auth/domain/email"
	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/email/mapper"
)

func (h *Handler) Generic(ctx context.Context, e *emailDomain.Email) error {
	updates := mapper.EmailDomainToUpdates(e)
	res := h.db.WithContext(ctx).Model(&rdbmodels.Email{}).
		Where(rdbmodels.EmailCols.ID+" = ?", e.ID().String()).
		Updates(updates)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return authErr.ErrEmailBindingNotFound
	}
	return nil
}
