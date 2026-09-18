package delete

import (
	"context"

	asseterrs "nfxidentity/errors/src/asset"
	rdbmodels "nfxidentity/modules/asset/infrastructure/rdb/models"

	"github.com/google/uuid"
)

func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	res := h.db.WithContext(ctx).Where(rdbmodels.FileCols.ID+" = ?", id.String()).Delete(&rdbmodels.File{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return asseterrs.ErrFileNotFound
	}
	return nil
}
