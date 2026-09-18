package update

import (
	"context"

	filesDomain "nfxidentity/modules/asset/domain/files"
	"nfxidentity/modules/asset/infrastructure/repository/files/mapper"
)

func (h *Handler) Generic(ctx context.Context, e *filesDomain.File) error {
	return h.db.WithContext(ctx).Save(mapper.FileDomainToModel(e)).Error
}
