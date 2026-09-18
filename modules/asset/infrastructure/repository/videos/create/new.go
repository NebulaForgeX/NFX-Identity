package create

import (
	"context"

	videosDomain "nfxidentity/modules/asset/domain/videos"
	"nfxidentity/modules/asset/infrastructure/repository/videos/mapper"
)

func (h *Handler) New(ctx context.Context, e *videosDomain.Video) error {
	return h.db.WithContext(ctx).Create(mapper.VideoDomainToModel(e)).Error
}
