package create

import (
	"context"

	identityDomain "nfxidentity/modules/auth/domain/identity"
	"nfxidentity/modules/auth/infrastructure/repository/identity/mapper"
)

func (h *Handler) New(ctx context.Context, i *identityDomain.Identity) error {
	return h.db.WithContext(ctx).Create(mapper.IdentityDomainToModel(i)).Error
}
