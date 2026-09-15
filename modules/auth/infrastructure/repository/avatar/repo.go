package avatar

import (
	"context"
	"time"

	"nfxidentity/modules/auth/domain/avatar"
	"nfxidentity/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repo struct{ db *gorm.DB }

func NewRepo(db *gorm.DB) *avatar.Repo {
	h := &repo{db: db}
	return &avatar.Repo{Create: h, Get: h, Update: h}
}

func (h *repo) table(kind string) (any, string) {
	if kind == "authority" {
		return &models.AuthorityProfileAvatar{}, "auth.AuthorityProfileAvatars"
	}
	return &models.ForgerProfileAvatar{}, "auth.ForgerProfileAvatars"
}

func (h *repo) New(ctx context.Context, a *avatar.Avatar) error {
	st := a.State()
	base := models.ProfileAvatar{ID: st.ID, ProfileID: st.ProfileID, ImageID: st.ImageID, IsActive: st.IsActive, CreatedAt: st.CreatedAt, UpdatedAt: st.UpdatedAt, DeletedAt: st.DeletedAt}
	if st.Kind == "authority" {
		return h.db.WithContext(ctx).Create(&models.AuthorityProfileAvatar{ProfileAvatar: base}).Error
	}
	return h.db.WithContext(ctx).Create(&models.ForgerProfileAvatar{ProfileAvatar: base}).Error
}

func (h *repo) ByProfileID(ctx context.Context, kind string, profileID uuid.UUID) ([]*avatar.Avatar, error) {
	var rows []models.ProfileAvatar
	tbl := "auth.ForgerProfileAvatars"
	if kind == "authority" {
		tbl = "auth.AuthorityProfileAvatars"
	}
	if err := h.db.WithContext(ctx).Table(tbl).Where("profile_id = ? AND deleted_at IS NULL", profileID).Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]*avatar.Avatar, 0, len(rows))
	for _, r := range rows {
		out = append(out, avatar.NewFromState(avatar.State{ID: r.ID, ProfileID: r.ProfileID, ImageID: r.ImageID, IsActive: r.IsActive, Kind: kind, CreatedAt: r.CreatedAt, UpdatedAt: r.UpdatedAt, DeletedAt: r.DeletedAt}))
	}
	return out, nil
}

func (h *repo) DeactivateActive(ctx context.Context, kind string, profileID uuid.UUID) error {
	model, _ := h.table(kind)
	return h.db.WithContext(ctx).Model(model).Where("profile_id = ? AND is_active = true", profileID).Updates(map[string]any{"is_active": false, "updated_at": time.Now()}).Error
}
