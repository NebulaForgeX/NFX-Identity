package get

import (
	"context"
	"time"

	"nfxidentity/constants"
	imagesDomain "nfxidentity/modules/asset/domain/images"
	rdbmodels "nfxidentity/modules/asset/infrastructure/rdb/models"
	"nfxidentity/modules/asset/infrastructure/repository/images/mapper"
)

func (h *Handler) ListStaleTmpBefore(ctx context.Context, cutoff time.Time, limit int) ([]*imagesDomain.Image, error) {
	var rows []rdbmodels.Image
	if err := h.db.WithContext(ctx).
		Where(rdbmodels.ImageCols.FilePath+" LIKE ?", constants.StaleTmpPathLike).
		Where(rdbmodels.ImageCols.CreatedAt+" < ?", cutoff).
		Order(rdbmodels.ImageCols.CreatedAt + " ASC").
		Limit(limit).
		Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]*imagesDomain.Image, 0, len(rows))
	for i := range rows {
		out = append(out, mapper.ImageModelToDomain(&rows[i]))
	}
	return out, nil
}
