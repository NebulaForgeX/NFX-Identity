package repository

import (
	"context"
	"errors"
	"nfxid/modules/auth/domain/user_role"
	userRoleDomainErrors "nfxid/modules/auth/domain/user_role/errors"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRolePGRepo struct {
	db *gorm.DB
}

func NewUserRolePGRepo(db *gorm.DB) *userRolePGRepo {
	return &userRolePGRepo{db: db}
}

func (r *userRolePGRepo) Create(ctx context.Context, ur *user_role.UserRole) error {
	m := mapper.UserRoleDomainToModel(ur)
	return r.db.WithContext(ctx).Create(&m).Error
}

func (r *userRolePGRepo) GetByID(ctx context.Context, id uuid.UUID) (*user_role.UserRole, error) {
	var m models.UserRole
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, userRoleDomainErrors.ErrUserRoleNotFound
		}
		return nil, err
	}
	return mapper.UserRoleModelToDomain(&m), nil
}

func (r *userRolePGRepo) GetByUserID(ctx context.Context, userID uuid.UUID) ([]*user_role.UserRole, error) {
	var models []models.UserRole
	if err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&models).Error; err != nil {
		return nil, err
	}

	entities := make([]*user_role.UserRole, len(models))
	for i := range models {
		entities[i] = mapper.UserRoleModelToDomain(&models[i])
	}
	return entities, nil
}

func (r *userRolePGRepo) GetByRoleID(ctx context.Context, roleID uuid.UUID) ([]*user_role.UserRole, error) {
	var models []models.UserRole
	if err := r.db.WithContext(ctx).
		Where("role_id = ?", roleID).
		Order("created_at DESC").
		Find(&models).Error; err != nil {
		return nil, err
	}

	entities := make([]*user_role.UserRole, len(models))
	for i := range models {
		entities[i] = mapper.UserRoleModelToDomain(&models[i])
	}
	return entities, nil
}

func (r *userRolePGRepo) Exists(ctx context.Context, userID, roleID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.UserRole{}).
		Where("user_id = ? AND role_id = ?", userID, roleID).
		Count(&count).Error
	return count > 0, err
}

func (r *userRolePGRepo) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.UserRole{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return userRoleDomainErrors.ErrUserRoleNotFound
	}
	return nil
}

func (r *userRolePGRepo) DeleteByUserAndRole(ctx context.Context, userID, roleID uuid.UUID) error {
	result := r.db.WithContext(ctx).
		Where("user_id = ? AND role_id = ?", userID, roleID).
		Delete(&models.UserRole{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return userRoleDomainErrors.ErrUserRoleNotFound
	}
	return nil
}
