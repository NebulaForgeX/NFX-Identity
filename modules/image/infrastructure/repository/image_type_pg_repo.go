package repository

import (
	"context"
	"errors"
	imageTypeDomain "nebulaid/modules/image/domain/image_type"
	imageTypeDomainErrors "nebulaid/modules/image/domain/image_type/errors"
	"nebulaid/modules/image/infrastructure/rdb/models"
	"nebulaid/modules/image/infrastructure/repository/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type imageTypePGRepo struct {
	db *gorm.DB
}

func NewImageTypePGRepo(db *gorm.DB) *imageTypePGRepo {
	return &imageTypePGRepo{db: db}
}

func (r *imageTypePGRepo) Create(ctx context.Context, it *imageTypeDomain.ImageType) error {
	m := mapper.ImageTypeDomainToModel(it)
	return r.db.WithContext(ctx).Create(&m).Error
}

func (r *imageTypePGRepo) Update(ctx context.Context, it *imageTypeDomain.ImageType) error {
	m := mapper.ImageTypeDomainToModel(it)
	updates := mapper.ImageTypeModelsToUpdates(m)
	return r.db.WithContext(ctx).
		Model(&models.ImageType{}).
		Where("id = ?", it.ID()).
		Updates(updates).Error
}

func (r *imageTypePGRepo) GetByID(ctx context.Context, id uuid.UUID) (*imageTypeDomain.ImageType, error) {
	var m models.ImageType
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, imageTypeDomainErrors.ErrImageTypeNotFound
		}
		return nil, err
	}
	return mapper.ImageTypeModelToDomain(&m), nil
}

func (r *imageTypePGRepo) GetByKey(ctx context.Context, key string) (*imageTypeDomain.ImageType, error) {
	var m models.ImageType
	if err := r.db.WithContext(ctx).Where("key = ?", key).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, imageTypeDomainErrors.ErrImageTypeNotFound
		}
		return nil, err
	}
	return mapper.ImageTypeModelToDomain(&m), nil
}

func (r *imageTypePGRepo) Exists(ctx context.Context, id uuid.UUID) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.ImageType{}).
		Where("id = ?", id).
		Count(&count).Error
	return count > 0, err
}

func (r *imageTypePGRepo) ExistsByKey(ctx context.Context, key string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.ImageType{}).
		Where("key = ?", key).
		Count(&count).Error
	return count > 0, err
}

func (r *imageTypePGRepo) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.ImageType{}, "id = ?", id).Error
}
