package update

import (
	"context"
	"nfxidentity/modules/auth/domain/refreshtoken"
	"nfxidentity/modules/auth/infrastructure/repository/refreshtoken/mapper"
)

func (h *Handler) Generic(ctx context.Context, t *refreshtoken.RefreshToken) error {
	return h.db.WithContext(ctx).Save(mapper.ToModel(t)).Error
}
