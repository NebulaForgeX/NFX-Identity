package get

import (
	"context"
	"errors"

	asseterrs "nfxidentity/errors/src/asset"
	filesDomain "nfxidentity/modules/asset/domain/files"
	rdbmodels "nfxidentity/modules/asset/infrastructure/rdb/models"
	"nfxidentity/modules/asset/infrastructure/repository/files/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*filesDomain.File, error) {
	var m rdbmodels.File
	if err := h.db.WithContext(ctx).Where(rdbmodels.FileCols.ID+" = ?", id.String()).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, asseterrs.ErrFileNotFound
		}
		return nil, err
	}
	return mapper.FileModelToDomain(&m), nil
}
