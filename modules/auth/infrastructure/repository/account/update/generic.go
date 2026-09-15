package update

import (
	"context"
	"nfxidentity/modules/auth/domain/account"
	"nfxidentity/modules/auth/infrastructure/repository/account/mapper"
)

func (h *Handler) Generic(ctx context.Context, a *account.Account) error {
	return h.db.WithContext(ctx).Save(mapper.ToModel(a)).Error
}
