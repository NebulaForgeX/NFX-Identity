package get

import (
	"context"
	"time"

	"nfxidentity/constants"
	videosDomain "nfxidentity/modules/asset/domain/videos"
	rdbmodels "nfxidentity/modules/asset/infrastructure/rdb/models"
	"nfxidentity/modules/asset/infrastructure/repository/videos/mapper"
)

func (h *Handler) ListStaleTmpBefore(ctx context.Context, cutoff time.Time, limit int) ([]*videosDomain.Video, error) {
	var rows []rdbmodels.Video
	if err := h.db.WithContext(ctx).
		Where(rdbmodels.VideoCols.FilePath+" LIKE ?", constants.StaleTmpPathLike).
		Where(rdbmodels.VideoCols.CreatedAt+" < ?", cutoff).
		Order(rdbmodels.VideoCols.CreatedAt + " ASC").
		Limit(limit).
		Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]*videosDomain.Video, 0, len(rows))
	for i := range rows {
		out = append(out, mapper.VideoModelToDomain(&rows[i]))
	}
	return out, nil
}
