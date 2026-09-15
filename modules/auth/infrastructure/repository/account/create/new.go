package create

import (
	"context"
	"nfxidentity/modules/auth/domain/account"
	"nfxidentity/modules/auth/infrastructure/repository/account/mapper"
)

func (h *Handler) New(ctx context.Context, a *account.Account) error {
	return h.db.WithContext(ctx).Create(mapper.ToModel(a)).Error
}
