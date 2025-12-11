package repository

import (
	"context"
	"errors"
	"nfxid/modules/auth/domain/profile"
	profileDomainErrors "nfxid/modules/auth/domain/profile/errors"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type profilePGRepo struct {
	db *gorm.DB
}

func NewProfilePGRepo(db *gorm.DB) *profilePGRepo {
	return &profilePGRepo{db: db}
}

func (r *profilePGRepo) Create(ctx context.Context, p *profile.Profile) error {
	m := mapper.ProfileDomainToModel(p)
	return r.db.WithContext(ctx).Create(&m).Error
}

func (r *profilePGRepo) Update(ctx context.Context, p *profile.Profile) error {
	m := mapper.ProfileDomainToModel(p)
	updates := mapper.ProfileModelsToUpdates(m)
	return r.db.WithContext(ctx).
		Model(&models.Profile{}).
		Where("id = ?", p.ID()).
		Updates(updates).Error
}

func (r *profilePGRepo) GetByID(ctx context.Context, id uuid.UUID) (*profile.Profile, error) {
	var m models.Profile
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, profileDomainErrors.ErrProfileNotFound
		}
		return nil, err
	}
	return mapper.ProfileModelToDomain(&m), nil
}

func (r *profilePGRepo) GetByUserID(ctx context.Context, userID uuid.UUID) (*profile.Profile, error) {
	var m models.Profile
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, profileDomainErrors.ErrProfileNotFound
		}
		return nil, err
	}
	return mapper.ProfileModelToDomain(&m), nil
}

func (r *profilePGRepo) Exists(ctx context.Context, id uuid.UUID) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.Profile{}).
		Where("id = ?", id).
		Count(&count).Error
	return count > 0, err
}

func (r *profilePGRepo) ExistsByUserID(ctx context.Context, userID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.Profile{}).
		Where("user_id = ?", userID).
		Count(&count).Error
	return count > 0, err
}

func (r *profilePGRepo) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.Profile{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return profileDomainErrors.ErrProfileNotFound
	}
	return nil
}
