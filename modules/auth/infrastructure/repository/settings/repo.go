package settings

import (
	"context"
	"errors"

	"nfxidentity/errors/src/auth"
	"nfxidentity/modules/auth/domain/settings"
	"nfxidentity/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repo struct{ db *gorm.DB }

func NewRepo(db *gorm.DB) *settings.Repo {
	h := &repo{db: db}
	return &settings.Repo{Create: h, Get: h, Update: h}
}

func (h *repo) New(ctx context.Context, s *settings.Settings) error {
	st := s.State()
	base := models.ProfileSettings{
		ID: st.ID, LoginNotification: st.LoginNotification,
		CreatedAt: st.CreatedAt, UpdatedAt: st.UpdatedAt, DeletedAt: st.DeletedAt,
	}
	if st.Kind == "authority" {
		return h.db.WithContext(ctx).Create(&models.AuthorityProfileSettings{ProfileSettings: base}).Error
	}
	return h.db.WithContext(ctx).Create(&models.ForgerProfileSettings{ProfileSettings: base}).Error
}

func (h *repo) ByID(ctx context.Context, kind string, id uuid.UUID) (*settings.Settings, error) {
	var m models.ProfileSettings
	q := h.db.WithContext(ctx).Table("auth.ForgerProfileSettings")
	if kind == "authority" {
		q = h.db.WithContext(ctx).Table("auth.AuthorityProfileSettings")
	}
	if err := q.Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if kind == "authority" {
				return nil, auth.ErrAuthorityProfileSettingsNotFound
			}
			return nil, auth.ErrForgerProfileSettingsNotFound
		}
		return nil, err
	}
	return settings.NewFromState(settings.State{
		ID: m.ID, Kind: kind, LoginNotification: m.LoginNotification,
		CreatedAt: m.CreatedAt, UpdatedAt: m.UpdatedAt, DeletedAt: m.DeletedAt,
	}), nil
}

func (h *repo) Generic(ctx context.Context, s *settings.Settings) error {
	st := s.State()
	base := models.ProfileSettings{
		ID: st.ID, LoginNotification: st.LoginNotification,
		CreatedAt: st.CreatedAt, UpdatedAt: st.UpdatedAt, DeletedAt: st.DeletedAt,
	}
	if st.Kind == "authority" {
		return h.db.WithContext(ctx).Save(&models.AuthorityProfileSettings{ProfileSettings: base}).Error
	}
	return h.db.WithContext(ctx).Save(&models.ForgerProfileSettings{ProfileSettings: base}).Error
}
