package create

import (
	"context"

	audiosDomain "nfxidentity/modules/asset/domain/audios"
	"nfxidentity/modules/asset/infrastructure/repository/audios/mapper"
)

func (h *Handler) New(ctx context.Context, e *audiosDomain.Audio) error {
	return h.db.WithContext(ctx).Create(mapper.AudioDomainToModel(e)).Error
}
