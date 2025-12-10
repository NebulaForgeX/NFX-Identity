package repository

import (
	"context"
	"errors"
	"nebulaid/modules/auth/domain/profile_badge"
	profileBadgeDomainErrors "nebulaid/modules/auth/domain/profile_badge/errors"
	"nebulaid/modules/auth/infrastructure/rdb/models"
	"nebulaid/modules/auth/infrastructure/repository/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type profileBadgePGRepo struct {
	db *gorm.DB
}

func NewProfileBadgePGRepo(db *gorm.DB) *profileBadgePGRepo {
	return &profileBadgePGRepo{db: db}
}

func (r *profileBadgePGRepo) Create(ctx context.Context, pb *profile_badge.ProfileBadge) error {
	m := mapper.ProfileBadgeDomainToModel(pb)
	return r.db.WithContext(ctx).Create(&m).Error
}

func (r *profileBadgePGRepo) Update(ctx context.Context, pb *profile_badge.ProfileBadge) error {
	m := mapper.ProfileBadgeDomainToModel(pb)
	updates := mapper.ProfileBadgeModelsToUpdates(m)
	return r.db.WithContext(ctx).
		Model(&models.ProfileBadge{}).
		Where("id = ?", pb.ID()).
		Updates(updates).Error
}

func (r *profileBadgePGRepo) GetByID(ctx context.Context, id uuid.UUID) (*profile_badge.ProfileBadge, error) {
	var m models.ProfileBadge
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, profileBadgeDomainErrors.ErrProfileBadgeNotFound
		}
		return nil, err
	}
	return mapper.ProfileBadgeModelToDomain(&m), nil
}

func (r *profileBadgePGRepo) GetByProfileID(ctx context.Context, profileID uuid.UUID) ([]*profile_badge.ProfileBadge, error) {
	var models []models.ProfileBadge
	if err := r.db.WithContext(ctx).
		Where("profile_id = ?", profileID).
		Order("earned_at DESC").
		Find(&models).Error; err != nil {
		return nil, err
	}

	entities := make([]*profile_badge.ProfileBadge, len(models))
	for i := range models {
		entities[i] = mapper.ProfileBadgeModelToDomain(&models[i])
	}
	return entities, nil
}

func (r *profileBadgePGRepo) GetByBadgeID(ctx context.Context, badgeID uuid.UUID) ([]*profile_badge.ProfileBadge, error) {
	var models []models.ProfileBadge
	if err := r.db.WithContext(ctx).
		Where("badge_id = ?", badgeID).
		Order("earned_at DESC").
		Find(&models).Error; err != nil {
		return nil, err
	}

	entities := make([]*profile_badge.ProfileBadge, len(models))
	for i := range models {
		entities[i] = mapper.ProfileBadgeModelToDomain(&models[i])
	}
	return entities, nil
}

func (r *profileBadgePGRepo) Exists(ctx context.Context, profileID, badgeID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.ProfileBadge{}).
		Where("profile_id = ? AND badge_id = ?", profileID, badgeID).
		Count(&count).Error
	return count > 0, err
}

func (r *profileBadgePGRepo) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.ProfileBadge{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return profileBadgeDomainErrors.ErrProfileBadgeNotFound
	}
	return nil
}

func (r *profileBadgePGRepo) DeleteByProfileAndBadge(ctx context.Context, profileID, badgeID uuid.UUID) error {
	result := r.db.WithContext(ctx).
		Where("profile_id = ? AND badge_id = ?", profileID, badgeID).
		Delete(&models.ProfileBadge{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return profileBadgeDomainErrors.ErrProfileBadgeNotFound
	}
	return nil
}
