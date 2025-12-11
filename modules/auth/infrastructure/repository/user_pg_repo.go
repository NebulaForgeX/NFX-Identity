package repository

import (
	"context"
	"errors"
	"nfxid/modules/auth/domain/user"
	userDomainErrors "nfxid/modules/auth/domain/user/errors"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userPGRepo struct {
	db *gorm.DB
}

func NewUserPGRepo(db *gorm.DB) *userPGRepo {
	return &userPGRepo{db: db}
}

func (r *userPGRepo) Create(ctx context.Context, u *user.User) error {
	m := mapper.UserDomainToModel(u)
	return r.db.WithContext(ctx).Create(&m).Error
}

func (r *userPGRepo) Update(ctx context.Context, u *user.User) error {
	m := mapper.UserDomainToModel(u)
	updates := mapper.UserModelsToUpdates(m)
	return r.db.WithContext(ctx).
		Model(&models.User{}).
		Where("id = ?", u.ID()).
		Updates(updates).Error
}

func (r *userPGRepo) GetByID(ctx context.Context, id uuid.UUID) (*user.User, error) {
	var m models.User
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, userDomainErrors.ErrUserNotFound
		}
		return nil, err
	}
	return mapper.UserModelToDomain(&m), nil
}

func (r *userPGRepo) GetByUsername(ctx context.Context, username string) (*user.User, error) {
	var m models.User
	if err := r.db.WithContext(ctx).Where("username = ?", username).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, userDomainErrors.ErrUserNotFound
		}
		return nil, err
	}
	return mapper.UserModelToDomain(&m), nil
}

func (r *userPGRepo) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	var m models.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, userDomainErrors.ErrUserNotFound
		}
		return nil, err
	}
	return mapper.UserModelToDomain(&m), nil
}

func (r *userPGRepo) GetByPhone(ctx context.Context, phone string) (*user.User, error) {
	var m models.User
	if err := r.db.WithContext(ctx).Where("phone = ?", phone).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, userDomainErrors.ErrUserNotFound
		}
		return nil, err
	}
	return mapper.UserModelToDomain(&m), nil
}

func (r *userPGRepo) Exists(ctx context.Context, id uuid.UUID) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.User{}).
		Where("id = ?", id).
		Count(&count).Error
	return count > 0, err
}

func (r *userPGRepo) ExistsByUsername(ctx context.Context, username string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.User{}).
		Where("username = ?", username).
		Count(&count).Error
	return count > 0, err
}

func (r *userPGRepo) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.User{}).
		Where("email = ?", email).
		Count(&count).Error
	return count > 0, err
}

func (r *userPGRepo) ExistsByPhone(ctx context.Context, phone string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.User{}).
		Where("phone = ?", phone).
		Count(&count).Error
	return count > 0, err
}

func (r *userPGRepo) UpdatePassword(ctx context.Context, userID uuid.UUID, hashedPassword string) error {
	result := r.db.WithContext(ctx).
		Model(&models.User{}).
		Where("id = ?", userID).
		Update(models.UserCols.PasswordHash, hashedPassword)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return userDomainErrors.ErrUserNotFound
	}
	return nil
}

func (r *userPGRepo) UpdateStatus(ctx context.Context, userID uuid.UUID, status string) error {
	result := r.db.WithContext(ctx).
		Model(&models.User{}).
		Where("id = ?", userID).
		Update(models.UserCols.Status, status)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return userDomainErrors.ErrUserNotFound
	}
	return nil
}

func (r *userPGRepo) Delete(ctx context.Context, userID uuid.UUID) error {
	result := r.db.WithContext(ctx).
		Where("id = ?", userID).
		Delete(&models.User{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return userDomainErrors.ErrUserNotFound
	}
	return nil
}
