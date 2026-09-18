package get

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	profileDomain "nfxidentity/modules/auth/domain/profile"
	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/profile/mapper"
)

//* =============================== Authority avatar_by_profile_id =============================== !//
func (h *Handler) AuthorityAvatarByProfileID(ctx context.Context, profileID uuid.UUID) (*profileDomain.AuthorityProfileAvatar, error) {
	var m rdbmodels.AuthorityProfileAvatar
	if err := h.db.WithContext(ctx).
		Where(rdbmodels.AuthorityProfileAvatarCols.ProfileID+" = ?", profileID.String()).
		Order(rdbmodels.AuthorityProfileAvatarCols.IsActive + " DESC, " + rdbmodels.AuthorityProfileAvatarCols.CreatedAt + " ASC").
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return mapper.AuthorityProfileAvatarModelToDomain(&m), nil
}

//* =============================== Forger avatar_by_profile_id =============================== !//
func (h *Handler) ForgerAvatarByProfileID(ctx context.Context, profileID uuid.UUID) (*profileDomain.ForgerProfileAvatar, error) {
	var m rdbmodels.ForgerProfileAvatar
	if err := h.db.WithContext(ctx).
		Where(rdbmodels.ForgerProfileAvatarCols.ProfileID+" = ?", profileID.String()).
		Order(rdbmodels.ForgerProfileAvatarCols.IsActive + " DESC, " + rdbmodels.ForgerProfileAvatarCols.CreatedAt + " ASC").
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return mapper.ForgerProfileAvatarModelToDomain(&m), nil
}
