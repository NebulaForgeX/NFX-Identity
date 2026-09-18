package background

import (
	"context"

	"nfxidentity/modules/auth/domain/background"
	"nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/pkgs/utils/timex"

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
	deleted := timex.TimeToGormDeletedAt(st.DeletedAt)
	if st.Kind == "authority" {
		return h.db.WithContext(ctx).Create(&models.Authorityprofilebackground{
			ID: st.ID, ProfileID: st.ProfileID, ImageID: st.ImageID, SortOrder: st.SortOrder,
			CreatedAt: st.CreatedAt, UpdatedAt: st.UpdatedAt, DeletedAt: deleted,
		}).Error
	}
	return h.db.WithContext(ctx).Create(&models.Forgerprofilebackground{
		ID: st.ID, ProfileID: st.ProfileID, ImageID: st.ImageID, SortOrder: st.SortOrder,
		CreatedAt: st.CreatedAt, UpdatedAt: st.UpdatedAt, DeletedAt: deleted,
	}).Error
}

func (h *repo) ByProfileID(ctx context.Context, kind string, profileID uuid.UUID) ([]*background.Background, error) {
	if kind == "authority" {
		var rows []models.Authorityprofilebackground
		if err := h.db.WithContext(ctx).Where("profile_id = ? AND deleted_at IS NULL", profileID).Order("sort_order").Find(&rows).Error; err != nil {
			return nil, err
		}
		return mapBackgrounds(kind, rows), nil
	}
	var rows []models.Forgerprofilebackground
	if err := h.db.WithContext(ctx).Where("profile_id = ? AND deleted_at IS NULL", profileID).Order("sort_order").Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]*background.Background, 0, len(rows))
	for _, r := range rows {
		out = append(out, background.NewFromState(background.State{
			ID: r.ID, ProfileID: r.ProfileID, ImageID: r.ImageID, SortOrder: r.SortOrder,
			Kind: kind, CreatedAt: r.CreatedAt, UpdatedAt: r.UpdatedAt, DeletedAt: timex.GormDeletedAtToTime(r.DeletedAt),
		}))
	}
	return out, nil
}

func mapBackgrounds(kind string, rows []models.Authorityprofilebackground) []*background.Background {
	out := make([]*background.Background, 0, len(rows))
	for _, r := range rows {
		out = append(out, background.NewFromState(background.State{
			ID: r.ID, ProfileID: r.ProfileID, ImageID: r.ImageID, SortOrder: r.SortOrder,
			Kind: kind, CreatedAt: r.CreatedAt, UpdatedAt: r.UpdatedAt, DeletedAt: timex.GormDeletedAtToTime(r.DeletedAt),
		}))
	}
	return out
}

func (h *repo) AllByProfileID(ctx context.Context, kind string, profileID uuid.UUID) error {
	if kind == "authority" {
		return h.db.WithContext(ctx).Where("profile_id = ?", profileID).Delete(&models.Authorityprofilebackground{}).Error
	}
	return h.db.WithContext(ctx).Where("profile_id = ?", profileID).Delete(&models.Forgerprofilebackground{}).Error
}
