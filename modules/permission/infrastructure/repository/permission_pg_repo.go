package repository

import (
	"context"
	"errors"
	"nfxid/modules/permission/domain/permission"
	permissionDomainErrors "nfxid/modules/permission/domain/permission/errors"
	"nfxid/modules/permission/infrastructure/rdb/models"
	"nfxid/modules/permission/infrastructure/repository/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type permissionPGRepo struct {
	db *gorm.DB
}

func NewPermissionPGRepo(db *gorm.DB) *permissionPGRepo {
	return &permissionPGRepo{db: db}
}

func (r *permissionPGRepo) Create(ctx context.Context, p *permission.Permission) error {
	m := mapper.PermissionDomainToModel(p)
	return r.db.WithContext(ctx).Create(&m).Error
}

func (r *permissionPGRepo) Update(ctx context.Context, p *permission.Permission) error {
	m := mapper.PermissionDomainToModel(p)
	updates := mapper.PermissionModelsToUpdates(m)
	return r.db.WithContext(ctx).
		Model(&models.Permission{}).
		Where("id = ?", p.ID()).
		Updates(updates).Error
}

func (r *permissionPGRepo) GetByID(ctx context.Context, id uuid.UUID) (*permission.Permission, error) {
	var m models.Permission
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, permissionDomainErrors.ErrPermissionNotFound
		}
		return nil, err
	}
	return mapper.PermissionModelToDomain(&m), nil
}

func (r *permissionPGRepo) GetByTag(ctx context.Context, tag string) (*permission.Permission, error) {
	var m models.Permission
	if err := r.db.WithContext(ctx).Where("tag = ?", tag).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, permissionDomainErrors.ErrPermissionNotFound
		}
		return nil, err
	}
	return mapper.PermissionModelToDomain(&m), nil
}

func (r *permissionPGRepo) GetByTags(ctx context.Context, tags []string) ([]*permission.Permission, error) {
	var models []models.Permission
	if err := r.db.WithContext(ctx).
		Where("tag IN ?", tags).
		Where("deleted_at IS NULL").
		Find(&models).Error; err != nil {
		return nil, err
	}

	entities := make([]*permission.Permission, len(models))
	for i := range models {
		entities[i] = mapper.PermissionModelToDomain(&models[i])
	}
	return entities, nil
}

func (r *permissionPGRepo) GetByCategory(ctx context.Context, category string) ([]*permission.Permission, error) {
	var models []models.Permission
	if err := r.db.WithContext(ctx).
		Where("category = ?", category).
		Where("deleted_at IS NULL").
		Order("tag ASC").
		Find(&models).Error; err != nil {
		return nil, err
	}

	entities := make([]*permission.Permission, len(models))
	for i := range models {
		entities[i] = mapper.PermissionModelToDomain(&models[i])
	}
	return entities, nil
}

func (r *permissionPGRepo) Exists(ctx context.Context, id uuid.UUID) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.Permission{}).
		Where("id = ?", id).
		Count(&count).Error
	return count > 0, err
}

func (r *permissionPGRepo) ExistsByTag(ctx context.Context, tag string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.Permission{}).
		Where("tag = ?", tag).
		Count(&count).Error
	return count > 0, err
}

func (r *permissionPGRepo) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.Permission{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return permissionDomainErrors.ErrPermissionNotFound
	}
	return nil
}

