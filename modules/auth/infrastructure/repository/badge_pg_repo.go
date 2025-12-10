package repository

import (
	"context"
	"errors"
	"nebulaid/modules/auth/domain/badge"
	badgeDomainErrors "nebulaid/modules/auth/domain/badge/errors"
	"nebulaid/modules/auth/infrastructure/rdb/models"
	"nebulaid/modules/auth/infrastructure/repository/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type badgePGRepo struct {
	db *gorm.DB
}

func NewBadgePGRepo(db *gorm.DB) *badgePGRepo {
	return &badgePGRepo{db: db}
}

func (r *badgePGRepo) Create(ctx context.Context, b *badge.Badge) error {
	m := mapper.BadgeDomainToModel(b)
	return r.db.WithContext(ctx).Create(&m).Error
}

func (r *badgePGRepo) Update(ctx context.Context, b *badge.Badge) error {
	m := mapper.BadgeDomainToModel(b)
	updates := mapper.BadgeModelsToUpdates(m)
	return r.db.WithContext(ctx).
		Model(&models.Badge{}).
		Where("id = ?", b.ID()).
		Updates(updates).Error
}

func (r *badgePGRepo) GetByID(ctx context.Context, id uuid.UUID) (*badge.Badge, error) {
	var m models.Badge
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, badgeDomainErrors.ErrNotFound
		}
		return nil, err
	}
	return mapper.BadgeModelToDomain(&m), nil
}
