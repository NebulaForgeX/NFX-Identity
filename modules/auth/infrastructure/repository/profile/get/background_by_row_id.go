package get

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	authErr "nfxidentity/errors/src/auth"
	profileDomain "nfxidentity/modules/auth/domain/profile"
	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/profile/mapper"
)

//* =============================== Authority background_by_row_id =============================== !//
func (h *Handler) AuthorityBackgroundByRowID(ctx context.Context, rowID uuid.UUID) (*profileDomain.AuthorityProfileBackground, error) {
	var m rdbmodels.AuthorityProfileBackground
	if err := h.db.WithContext(ctx).Where(rdbmodels.AuthorityProfileBackgroundCols.ID+" = ?", rowID.String()).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, authErr.ErrAuthorityProfileNotFound
		}
		return nil, err
	}
	return mapper.AuthorityProfileBackgroundModelToDomain(&m), nil
}

//* =============================== Forger background_by_row_id =============================== !//
func (h *Handler) ForgerBackgroundByRowID(ctx context.Context, rowID uuid.UUID) (*profileDomain.ForgerProfileBackground, error) {
	var m rdbmodels.ForgerProfileBackground
	if err := h.db.WithContext(ctx).Where(rdbmodels.ForgerProfileBackgroundCols.ID+" = ?", rowID.String()).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, authErr.ErrForgerProfileNotFound
		}
		return nil, err
	}
	return mapper.ForgerProfileBackgroundModelToDomain(&m), nil
}
