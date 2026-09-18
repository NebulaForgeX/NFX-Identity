package get

import (
	"context"
	"time"

	"nfxidentity/constants"
	filesDomain "nfxidentity/modules/asset/domain/files"
	rdbmodels "nfxidentity/modules/asset/infrastructure/rdb/models"
	"nfxidentity/modules/asset/infrastructure/repository/files/mapper"
)

func (h *Handler) ListStaleTmpBefore(ctx context.Context, cutoff time.Time, limit int) ([]*filesDomain.File, error) {
	var rows []rdbmodels.File
	if err := h.db.WithContext(ctx).
		Where(rdbmodels.FileCols.FilePath+" LIKE ?", constants.StaleTmpPathLike).
		Where(rdbmodels.FileCols.CreatedAt+" < ?", cutoff).
		Order(rdbmodels.FileCols.CreatedAt + " ASC").
		Limit(limit).
		Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]*filesDomain.File, 0, len(rows))
	for i := range rows {
		out = append(out, mapper.FileModelToDomain(&rows[i]))
	}
	return out, nil
}
