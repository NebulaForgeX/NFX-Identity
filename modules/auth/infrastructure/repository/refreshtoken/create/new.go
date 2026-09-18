package create

import (
	"context"

	refreshtokenDomain "nfxidentity/modules/auth/domain/refreshtoken"
	"nfxidentity/modules/auth/infrastructure/repository/refreshtoken/mapper"
)

func (h *Handler) New(ctx context.Context, t *refreshtokenDomain.RefreshToken) error {
	return h.db.WithContext(ctx).Create(mapper.RefreshTokenDomainToModel(t)).Error
}
