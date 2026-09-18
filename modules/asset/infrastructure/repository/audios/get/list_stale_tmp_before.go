package get

import (
	"context"
	"time"

	"nfxidentity/constants"
	audiosDomain "nfxidentity/modules/asset/domain/audios"
	rdbmodels "nfxidentity/modules/asset/infrastructure/rdb/models"
	"nfxidentity/modules/asset/infrastructure/repository/audios/mapper"
)

func (h *Handler) ListStaleTmpBefore(ctx context.Context, cutoff time.Time, limit int) ([]*audiosDomain.Audio, error) {
	var rows []rdbmodels.Audio
	if err := h.db.WithContext(ctx).
		Where(rdbmodels.AudioCols.FilePath+" LIKE ?", constants.StaleTmpPathLike).
		Where(rdbmodels.AudioCols.CreatedAt+" < ?", cutoff).
		Order(rdbmodels.AudioCols.CreatedAt + " ASC").
		Limit(limit).
		Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]*audiosDomain.Audio, 0, len(rows))
	for i := range rows {
		out = append(out, mapper.AudioModelToDomain(&rows[i]))
	}
	return out, nil
}
