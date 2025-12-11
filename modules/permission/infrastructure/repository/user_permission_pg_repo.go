package repository

import (
	"context"
	"errors"
	"nfxid/modules/permission/domain/user_permission"
	userPermissionDomainErrors "nfxid/modules/permission/domain/user_permission/errors"
	"nfxid/modules/permission/infrastructure/rdb/models"
	"nfxid/modules/permission/infrastructure/repository/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userPermissionPGRepo struct {
	db *gorm.DB
}

func NewUserPermissionPGRepo(db *gorm.DB) *userPermissionPGRepo {
	return &userPermissionPGRepo{db: db}
}

func (r *userPermissionPGRepo) Create(ctx context.Context, up *user_permission.UserPermission) error {
	m := mapper.UserPermissionDomainToModel(up)
	return r.db.WithContext(ctx).Create(&m).Error
}

func (r *userPermissionPGRepo) GetByUserID(ctx context.Context, userID uuid.UUID) ([]*user_permission.UserPermission, error) {
	var models []models.UserPermission
	if err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Where("deleted_at IS NULL").
		Order("created_at DESC").
		Find(&models).Error; err != nil {
		return nil, err
	}

	entities := make([]*user_permission.UserPermission, len(models))
	for i := range models {
		entities[i] = mapper.UserPermissionModelToDomain(&models[i])
	}
	return entities, nil
}

func (r *userPermissionPGRepo) GetByUserIDAndPermissionID(ctx context.Context, userID, permissionID uuid.UUID) (*user_permission.UserPermission, error) {
	var m models.UserPermission
	if err := r.db.WithContext(ctx).
		Where("user_id = ? AND permission_id = ?", userID, permissionID).
		Where("deleted_at IS NULL").
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, userPermissionDomainErrors.ErrUserPermissionNotFound
		}
		return nil, err
	}
	return mapper.UserPermissionModelToDomain(&m), nil
}

func (r *userPermissionPGRepo) GetPermissionTagsByUserID(ctx context.Context, userID uuid.UUID) ([]string, error) {
	var tags []string
	err := r.db.WithContext(ctx).
		Model(&models.UserPermission{}).
		Select("p.tag").
		Joins("JOIN permission.permissions p ON p.id = user_permissions.permission_id").
		Where("user_permissions.user_id = ?", userID).
		Where("user_permissions.deleted_at IS NULL").
		Where("p.deleted_at IS NULL").
		Pluck("p.tag", &tags).Error
	return tags, err
}

func (r *userPermissionPGRepo) Exists(ctx context.Context, userID, permissionID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.UserPermission{}).
		Where("user_id = ? AND permission_id = ?", userID, permissionID).
		Where("deleted_at IS NULL").
		Count(&count).Error
	return count > 0, err
}

func (r *userPermissionPGRepo) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.UserPermission{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return userPermissionDomainErrors.ErrUserPermissionNotFound
	}
	return nil
}

func (r *userPermissionPGRepo) DeleteByUserIDAndPermissionID(ctx context.Context, userID, permissionID uuid.UUID) error {
	result := r.db.WithContext(ctx).
		Where("user_id = ? AND permission_id = ?", userID, permissionID).
		Delete(&models.UserPermission{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return userPermissionDomainErrors.ErrUserPermissionNotFound
	}
	return nil
}

func (r *userPermissionPGRepo) DeleteByUserID(ctx context.Context, userID uuid.UUID) error {
	result := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Delete(&models.UserPermission{})

	if result.Error != nil {
		return result.Error
	}
	return nil
}

