package update

import (
	"context"

	audiosDomain "nfxidentity/modules/asset/domain/audios"
	"nfxidentity/modules/asset/infrastructure/repository/audios/mapper"
)

func (h *Handler) Generic(ctx context.Context, e *audiosDomain.Audio) error {
	return h.db.WithContext(ctx).Save(mapper.AudioDomainToModel(e)).Error
}
