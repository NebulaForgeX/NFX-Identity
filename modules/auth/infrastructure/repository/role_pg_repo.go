package repository

import (
	"context"
	"errors"
	"nebulaid/modules/auth/domain/role"
	roleDomainErrors "nebulaid/modules/auth/domain/role/errors"
	"nebulaid/modules/auth/infrastructure/rdb/models"
	"nebulaid/modules/auth/infrastructure/repository/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type rolePGRepo struct {
	db *gorm.DB
}

func NewRolePGRepo(db *gorm.DB) *rolePGRepo {
	return &rolePGRepo{db: db}
}

func (r *rolePGRepo) Create(ctx context.Context, ro *role.Role) error {
	m := mapper.RoleDomainToModel(ro)
	return r.db.WithContext(ctx).Create(&m).Error
}

func (r *rolePGRepo) Update(ctx context.Context, ro *role.Role) error {
	m := mapper.RoleDomainToModel(ro)
	updates := mapper.RoleModelsToUpdates(m)
	return r.db.WithContext(ctx).
		Model(&models.Role{}).
		Where("id = ?", ro.ID()).
		Updates(updates).Error
}

func (r *rolePGRepo) GetByID(ctx context.Context, id uuid.UUID) (*role.Role, error) {
	var m models.Role
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, roleDomainErrors.ErrRoleNotFound
		}
		return nil, err
	}
	return mapper.RoleModelToDomain(&m), nil
}

func (r *rolePGRepo) GetByName(ctx context.Context, name string) (*role.Role, error) {
	var m models.Role
	if err := r.db.WithContext(ctx).Where("name = ?", name).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, roleDomainErrors.ErrRoleNotFound
		}
		return nil, err
	}
	return mapper.RoleModelToDomain(&m), nil
}

func (r *rolePGRepo) Exists(ctx context.Context, id uuid.UUID) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.Role{}).
		Where("id = ?", id).
		Count(&count).Error
	return count > 0, err
}

func (r *rolePGRepo) ExistsByName(ctx context.Context, name string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.Role{}).
		Where("name = ?", name).
		Count(&count).Error
	return count > 0, err
}

func (r *rolePGRepo) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.Role{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return roleDomainErrors.ErrRoleNotFound
	}
	return nil
}
