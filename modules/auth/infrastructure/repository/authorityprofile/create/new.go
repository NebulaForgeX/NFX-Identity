package create

import (
	"context"
	"nfxidentity/modules/auth/domain/authorityprofile"
	"nfxidentity/modules/auth/infrastructure/repository/authorityprofile/mapper"
)

func (h *Handler) New(ctx context.Context, p *authorityprofile.Profile) error {
	return h.db.WithContext(ctx).Create(mapper.ToModel(p)).Error
}
