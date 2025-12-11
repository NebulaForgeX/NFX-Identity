package repository

import (
	"context"
	"errors"
	"nfxid/modules/auth/domain/profile_occupation"
	occupationDomainErrors "nfxid/modules/auth/domain/profile_occupation/errors"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type occupationPGRepo struct {
	db *gorm.DB
}

func NewOccupationPGRepo(db *gorm.DB) *occupationPGRepo {
	return &occupationPGRepo{db: db}
}

func (r *occupationPGRepo) Create(ctx context.Context, o *occupation.Occupation) error {
	m := mapper.OccupationDomainToModel(o)
	return r.db.WithContext(ctx).Create(&m).Error
}

func (r *occupationPGRepo) Update(ctx context.Context, o *occupation.Occupation) error {
	m := mapper.OccupationDomainToModel(o)
	updates := mapper.OccupationModelsToUpdates(m)
	return r.db.WithContext(ctx).
		Model(&models.Occupation{}).
		Where("id = ?", o.ID()).
		Updates(updates).Error
}

func (r *occupationPGRepo) GetByID(ctx context.Context, id uuid.UUID) (*occupation.Occupation, error) {
	var m models.Occupation
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, occupationDomainErrors.ErrOccupationNotFound
		}
		return nil, err
	}
	return mapper.OccupationModelToDomain(&m), nil
}

func (r *occupationPGRepo) GetByProfileID(ctx context.Context, profileID uuid.UUID) ([]*occupation.Occupation, error) {
	var models []models.Occupation
	if err := r.db.WithContext(ctx).
		Where("profile_id = ?", profileID).
		Where("deleted_at IS NULL").
		Order("start_date DESC").
		Find(&models).Error; err != nil {
		return nil, err
	}

	entities := make([]*occupation.Occupation, len(models))
	for i := range models {
		entities[i] = mapper.OccupationModelToDomain(&models[i])
	}
	return entities, nil
}

func (r *occupationPGRepo) Exists(ctx context.Context, id uuid.UUID) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.Occupation{}).
		Where("id = ?", id).
		Count(&count).Error
	return count > 0, err
}

func (r *occupationPGRepo) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.Occupation{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return occupationDomainErrors.ErrOccupationNotFound
	}
	return nil
}
