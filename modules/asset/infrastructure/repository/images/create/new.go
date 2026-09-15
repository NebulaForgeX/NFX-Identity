package create

import (
	"context"

	"nfxidentity/modules/asset/domain/images"
	"nfxidentity/modules/asset/infrastructure/repository/images/mapper"
)

func (h *Handler) New(ctx context.Context, img *images.Image) error {
	return h.db.WithContext(ctx).Create(mapper.ToModel(img)).Error
}
