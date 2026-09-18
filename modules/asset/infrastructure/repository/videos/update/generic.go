package update

import (
	"context"

	videosDomain "nfxidentity/modules/asset/domain/videos"
	"nfxidentity/modules/asset/infrastructure/repository/videos/mapper"
)

func (h *Handler) Generic(ctx context.Context, e *videosDomain.Video) error {
	return h.db.WithContext(ctx).Save(mapper.VideoDomainToModel(e)).Error
}
