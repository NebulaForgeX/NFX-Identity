package update

import (
	"context"

	imagesDomain "nfxidentity/modules/asset/domain/images"
	"nfxidentity/modules/asset/infrastructure/repository/images/mapper"
)

func (h *Handler) Generic(ctx context.Context, e *imagesDomain.Image) error {
	return h.db.WithContext(ctx).Save(mapper.ImageDomainToModel(e)).Error
}
