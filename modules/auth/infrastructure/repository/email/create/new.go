package create

import (
	"context"

	emailDomain "nfxidentity/modules/auth/domain/email"
	"nfxidentity/modules/auth/infrastructure/repository/email/mapper"
)

func (h *Handler) New(ctx context.Context, e *emailDomain.Email) error {
	return h.db.WithContext(ctx).Create(mapper.EmailDomainToModel(e)).Error
}
