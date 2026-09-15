package background

import (
	"context"

	"nfxidentity/modules/auth/domain/background"
	"nfxidentity/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repo struct{ db *gorm.DB }

func NewRepo(db *gorm.DB) *background.Repo {
	h := &repo{db: db}
	return &background.Repo{Create: h, Get: h, Delete: h}
}

func (h *repo) New(ctx context.Context, b *background.Background) error {
	st := b.State()
	base := models.ProfileBackground{
		ID: st.ID, ProfileID: st.ProfileID, ImageID: st.ImageID, SortOrder: st.SortOrder,
		CreatedAt: st.CreatedAt, UpdatedAt: st.UpdatedAt, DeletedAt: st.DeletedAt,
	}
	if st.Kind == "authority" {
		return h.db.WithContext(ctx).Create(&models.AuthorityProfileBackground{ProfileBackground: base}).Error
	}
	return h.db.WithContext(ctx).Create(&models.ForgerProfileBackground{ProfileBackground: base}).Error
}

func (h *repo) ByProfileID(ctx context.Context, kind string, profileID uuid.UUID) ([]*background.Background, error) {
	tbl := "auth.ForgerProfileBackgrounds"
	if kind == "authority" {
		tbl = "auth.AuthorityProfileBackgrounds"
	}
	var rows []models.ProfileBackground
	if err := h.db.WithContext(ctx).Table(tbl).Where("profile_id = ? AND deleted_at IS NULL", profileID).Order("sort_order").Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]*background.Background, 0, len(rows))
	for _, r := range rows {
		out = append(out, background.NewFromState(background.State{
			ID: r.ID, ProfileID: r.ProfileID, ImageID: r.ImageID, SortOrder: r.SortOrder,
			Kind: kind, CreatedAt: r.CreatedAt, UpdatedAt: r.UpdatedAt, DeletedAt: r.DeletedAt,
		}))
	}
	return out, nil
}

func (h *repo) AllByProfileID(ctx context.Context, kind string, profileID uuid.UUID) error {
	if kind == "authority" {
		return h.db.WithContext(ctx).Where("profile_id = ?", profileID).Delete(&models.AuthorityProfileBackground{}).Error
	}
	return h.db.WithContext(ctx).Where("profile_id = ?", profileID).Delete(&models.ForgerProfileBackground{}).Error
}
