package repository

import (
	"context"
	"errors"
	"nebulaid/modules/auth/domain/education"
	educationDomainErrors "nebulaid/modules/auth/domain/education/errors"
	"nebulaid/modules/auth/infrastructure/rdb/models"
	"nebulaid/modules/auth/infrastructure/repository/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type educationPGRepo struct {
	db *gorm.DB
}

func NewEducationPGRepo(db *gorm.DB) *educationPGRepo {
	return &educationPGRepo{db: db}
}

func (r *educationPGRepo) Create(ctx context.Context, e *education.Education) error {
	m := mapper.EducationDomainToModel(e)
	return r.db.WithContext(ctx).Create(&m).Error
}

func (r *educationPGRepo) Update(ctx context.Context, e *education.Education) error {
	m := mapper.EducationDomainToModel(e)
	updates := mapper.EducationModelsToUpdates(m)
	return r.db.WithContext(ctx).
		Model(&models.Education{}).
		Where("id = ?", e.ID()).
		Updates(updates).Error
}

func (r *educationPGRepo) GetByID(ctx context.Context, id uuid.UUID) (*education.Education, error) {
	var m models.Education
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, educationDomainErrors.ErrEducationNotFound
		}
		return nil, err
	}
	return mapper.EducationModelToDomain(&m), nil
}

func (r *educationPGRepo) GetByProfileID(ctx context.Context, profileID uuid.UUID) ([]*education.Education, error) {
	var models []models.Education
	if err := r.db.WithContext(ctx).
		Where("profile_id = ?", profileID).
		Where("deleted_at IS NULL").
		Order("start_date DESC").
		Find(&models).Error; err != nil {
		return nil, err
	}

	entities := make([]*education.Education, len(models))
	for i := range models {
		entities[i] = mapper.EducationModelToDomain(&models[i])
	}
	return entities, nil
}

func (r *educationPGRepo) Exists(ctx context.Context, id uuid.UUID) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.Education{}).
		Where("id = ?", id).
		Count(&count).Error
	return count > 0, err
}

func (r *educationPGRepo) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.Education{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return educationDomainErrors.ErrEducationNotFound
	}
	return nil
}
