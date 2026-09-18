package create

import (
	"context"

	imagesDomain "nfxidentity/modules/asset/domain/images"
	"nfxidentity/modules/asset/infrastructure/repository/images/mapper"
)

func (h *Handler) New(ctx context.Context, e *imagesDomain.Image) error {
	return h.db.WithContext(ctx).Create(mapper.ImageDomainToModel(e)).Error
}
