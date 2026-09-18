package create

import (
	"context"
	"nfxidentity/modules/auth/domain/email"
	"nfxidentity/modules/auth/infrastructure/repository/email/mapper"
)

func (h *Handler) New(ctx context.Context, e *email.Email) error {
	return h.db.WithContext(ctx).Create(mapper.ToModel(e)).Error
}
