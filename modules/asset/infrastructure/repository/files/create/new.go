package create

import (
	"context"

	filesDomain "nfxidentity/modules/asset/domain/files"
	"nfxidentity/modules/asset/infrastructure/repository/files/mapper"
)

func (h *Handler) New(ctx context.Context, e *filesDomain.File) error {
	return h.db.WithContext(ctx).Create(mapper.FileDomainToModel(e)).Error
}
