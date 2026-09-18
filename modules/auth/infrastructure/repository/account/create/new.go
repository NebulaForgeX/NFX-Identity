package create

import (
	"context"

	accountDomain "nfxidentity/modules/auth/domain/account"
	"nfxidentity/modules/auth/infrastructure/repository/account/mapper"
)

func (h *Handler) New(ctx context.Context, u *accountDomain.Account) error {
	return h.db.WithContext(ctx).Create(mapper.AccountDomainToModel(u)).Error
}
