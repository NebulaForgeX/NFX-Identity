package settings

import (
	"context"
	"errors"

	"nfxidentity/errors/src/auth"
	"nfxidentity/modules/auth/domain/settings"
	"nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/pkgs/utils/timex"

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
	deleted := timex.TimeToGormDeletedAt(st.DeletedAt)
	if st.Kind == "authority" {
		return h.db.WithContext(ctx).Create(&models.Authorityprofilesetting{
			ID: st.ID, LoginNotification: st.LoginNotification,
			CreatedAt: st.CreatedAt, UpdatedAt: st.UpdatedAt, DeletedAt: deleted,
		}).Error
	}
	return h.db.WithContext(ctx).Create(&models.Forgerprofilesetting{
		ID: st.ID, LoginNotification: st.LoginNotification,
		CreatedAt: st.CreatedAt, UpdatedAt: st.UpdatedAt, DeletedAt: deleted,
	}).Error
}

func (h *repo) ByID(ctx context.Context, kind string, id uuid.UUID) (*settings.Settings, error) {
	if kind == "authority" {
		var m models.Authorityprofilesetting
		if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, auth.ErrAuthorityProfileSettingsNotFound
			}
			return nil, err
		}
		return settings.NewFromState(settings.State{
			ID: m.ID, Kind: kind, LoginNotification: m.LoginNotification,
			CreatedAt: m.CreatedAt, UpdatedAt: m.UpdatedAt, DeletedAt: timex.GormDeletedAtToTime(m.DeletedAt),
		}), nil
	}
	var m models.Forgerprofilesetting
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, auth.ErrForgerProfileSettingsNotFound
		}
		return nil, err
	}
	return settings.NewFromState(settings.State{
		ID: m.ID, Kind: kind, LoginNotification: m.LoginNotification,
		CreatedAt: m.CreatedAt, UpdatedAt: m.UpdatedAt, DeletedAt: timex.GormDeletedAtToTime(m.DeletedAt),
	}), nil
}

func (h *repo) Generic(ctx context.Context, s *settings.Settings) error {
	st := s.State()
	deleted := timex.TimeToGormDeletedAt(st.DeletedAt)
	if st.Kind == "authority" {
		return h.db.WithContext(ctx).Save(&models.Authorityprofilesetting{
			ID: st.ID, LoginNotification: st.LoginNotification,
			CreatedAt: st.CreatedAt, UpdatedAt: st.UpdatedAt, DeletedAt: deleted,
		}).Error
	}
	return h.db.WithContext(ctx).Save(&models.Forgerprofilesetting{
		ID: st.ID, LoginNotification: st.LoginNotification,
		CreatedAt: st.CreatedAt, UpdatedAt: st.UpdatedAt, DeletedAt: deleted,
	}).Error
}
