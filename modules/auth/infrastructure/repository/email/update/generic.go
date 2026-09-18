package update

import (
	"context"
	"nfxidentity/modules/auth/domain/email"
	"nfxidentity/modules/auth/infrastructure/repository/email/mapper"
)

func (h *Handler) Generic(ctx context.Context, e *email.Email) error {
	return h.db.WithContext(ctx).Save(mapper.ToModel(e)).Error
}
