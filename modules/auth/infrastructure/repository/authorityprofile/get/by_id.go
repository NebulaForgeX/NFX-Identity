package get

import (
	"context"
	"errors"

	"nfxidentity/errors/src/auth"
	"nfxidentity/modules/auth/domain/authorityprofile"
	"nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/authorityprofile/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*authorityprofile.Profile, error) {
	var m models.AuthorityProfile
	if err := h.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, auth.ErrAuthorityProfileNotFound
		}
		return nil, err
	}
	return mapper.ToDomain(&m), nil
}
