package update

import (
	"context"

	authErr "nfxidentity/errors/src/auth"
	accountDomain "nfxidentity/modules/auth/domain/account"
	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/account/mapper"
)

func (h *Handler) Generic(ctx context.Context, u *accountDomain.Account) error {
	updates := mapper.AccountDomainToUpdates(u)
	res := h.db.WithContext(ctx).Model(&rdbmodels.Account{}).
		Where(rdbmodels.AccountCols.ID+" = ?", u.ID().String()).
		Updates(updates)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return authErr.ErrAccountNotFound
	}
	return nil
}
