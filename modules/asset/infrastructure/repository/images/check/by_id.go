package check

import (
	"context"

	rdbmodels "nfxidentity/modules/asset/infrastructure/rdb/models"

	"github.com/google/uuid"
)

func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (bool, error) {
	var n int64
	if err := h.db.WithContext(ctx).Model(&rdbmodels.Image{}).
		Where(rdbmodels.ImageCols.ID+" = ?", id.String()).
		Count(&n).Error; err != nil {
		return false, err
	}
	return n > 0, nil
}
