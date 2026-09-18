package create

import (
	"context"

	phoneDomain "nfxidentity/modules/auth/domain/phone"
	"nfxidentity/modules/auth/infrastructure/repository/phone/mapper"
)

func (h *Handler) New(ctx context.Context, p *phoneDomain.Phone) error {
	return h.db.WithContext(ctx).Create(mapper.PhoneDomainToModel(p)).Error
}
