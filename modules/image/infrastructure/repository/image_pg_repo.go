package repository

import (
	"context"
	"errors"
	imageDomain "nfxid/modules/image/domain/image"
	imageDomainErrors "nfxid/modules/image/domain/image/errors"
	"nfxid/modules/image/infrastructure/rdb/models"
	"nfxid/modules/image/infrastructure/repository/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type imagePGRepo struct {
	db *gorm.DB
}

func NewImagePGRepo(db *gorm.DB) *imagePGRepo {
	return &imagePGRepo{db: db}
}

func (r *imagePGRepo) Create(ctx context.Context, img *imageDomain.Image) error {
	m := mapper.ImageDomainToModel(img)
	return r.db.WithContext(ctx).Create(&m).Error
}

func (r *imagePGRepo) Update(ctx context.Context, img *imageDomain.Image) error {
	m := mapper.ImageDomainToModel(img)
	updates := mapper.ImageModelsToUpdates(m)
	return r.db.WithContext(ctx).
		Model(&models.Image{}).
		Where("id = ?", img.ID()).
		Updates(updates).Error
}

func (r *imagePGRepo) GetByID(ctx context.Context, id uuid.UUID) (*imageDomain.Image, error) {
	var m models.Image
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, imageDomainErrors.ErrImageNotFound
		}
		return nil, err
	}
	return mapper.ImageModelToDomain(&m), nil
}

func (r *imagePGRepo) GetByFilename(ctx context.Context, filename string) (*imageDomain.Image, error) {
	var m models.Image
	if err := r.db.WithContext(ctx).Where("filename = ?", filename).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, imageDomainErrors.ErrImageNotFound
		}
		return nil, err
	}
	return mapper.ImageModelToDomain(&m), nil
}

func (r *imagePGRepo) GetByStoragePath(ctx context.Context, storagePath string) (*imageDomain.Image, error) {
	var m models.Image
	if err := r.db.WithContext(ctx).Where("storage_path = ?", storagePath).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, imageDomainErrors.ErrImageNotFound
		}
		return nil, err
	}
	return mapper.ImageModelToDomain(&m), nil
}

func (r *imagePGRepo) Exists(ctx context.Context, id uuid.UUID) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.Image{}).
		Where("id = ?", id).
		Count(&count).Error
	return count > 0, err
}

func (r *imagePGRepo) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.Image{}, "id = ?", id).Error
}

func (r *imagePGRepo) DeleteByStoragePath(ctx context.Context, storagePath string) error {
	return r.db.WithContext(ctx).Delete(&models.Image{}, "storage_path = ?", storagePath).Error
}
