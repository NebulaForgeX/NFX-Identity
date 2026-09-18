package create

import (
	"context"
	"nfxidentity/modules/auth/domain/refreshtoken"
	"nfxidentity/modules/auth/infrastructure/repository/refreshtoken/mapper"
)

func (h *Handler) New(ctx context.Context, t *refreshtoken.RefreshToken) error {
	return h.db.WithContext(ctx).Create(mapper.ToModel(t)).Error
}
