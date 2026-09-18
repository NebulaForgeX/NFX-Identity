package create

import (
	"context"
	"nfxidentity/modules/auth/domain/forgerprofile"
	"nfxidentity/modules/auth/infrastructure/repository/forgerprofile/mapper"
)

func (h *Handler) New(ctx context.Context, p *forgerprofile.Profile) error {
	return h.db.WithContext(ctx).Create(mapper.ToModel(p)).Error
}
