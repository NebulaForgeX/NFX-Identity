package update

import (
	"context"

	"nfxidentity/modules/asset/domain/images"
	"nfxidentity/modules/asset/infrastructure/repository/images/mapper"
)

func (h *Handler) Generic(ctx context.Context, img *images.Image) error {
	return h.db.WithContext(ctx).Save(mapper.ToModel(img)).Error
}
