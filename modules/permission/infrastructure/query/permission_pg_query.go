package query

import (
	"context"
	"errors"
	permissionAppViews "nfxid/modules/permission/application/permission/views"
	permissionDomainErrors "nfxid/modules/permission/domain/permission/errors"
	"nfxid/modules/permission/infrastructure/query/mapper"
	"nfxid/modules/permission/infrastructure/rdb/models"
	"nfxid/pkgs/utils/slice"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type permissionPGQuery struct {
	db *gorm.DB
}

func NewPermissionPGQuery(db *gorm.DB) *permissionPGQuery {
	return &permissionPGQuery{db: db}
}

func (q *permissionPGQuery) GetByID(ctx context.Context, id uuid.UUID) (*permissionAppViews.PermissionView, error) {
	var m models.Permission
	if err := q.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, permissionDomainErrors.ErrPermissionNotFound
		}
		return nil, err
	}
	view := mapper.PermissionModelToAppView(&m)
	return &view, nil
}

func (q *permissionPGQuery) GetByTag(ctx context.Context, tag string) (*permissionAppViews.PermissionView, error) {
	var m models.Permission
	if err := q.db.WithContext(ctx).Where("tag = ?", tag).Where("deleted_at IS NULL").First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, permissionDomainErrors.ErrPermissionNotFound
		}
		return nil, err
	}
	view := mapper.PermissionModelToAppView(&m)
	return &view, nil
}

func (q *permissionPGQuery) GetByTags(ctx context.Context, tags []string) ([]*permissionAppViews.PermissionView, error) {
	var items []models.Permission
	if err := q.db.WithContext(ctx).
		Where("tag IN ?", tags).
		Where("deleted_at IS NULL").
		Order("tag ASC").
		Find(&items).Error; err != nil {
		return nil, err
	}
	return slice.MapP(items, func(m *models.Permission) *permissionAppViews.PermissionView {
		view := mapper.PermissionModelToAppView(m)
		return &view
	}), nil
}

func (q *permissionPGQuery) GetByCategory(ctx context.Context, category string) ([]*permissionAppViews.PermissionView, error) {
	var items []models.Permission
	if err := q.db.WithContext(ctx).
		Where("category = ?", category).
		Where("deleted_at IS NULL").
		Order("tag ASC").
		Find(&items).Error; err != nil {
		return nil, err
	}
	return slice.MapP(items, func(m *models.Permission) *permissionAppViews.PermissionView {
		view := mapper.PermissionModelToAppView(m)
		return &view
	}), nil
}

func (q *permissionPGQuery) List(ctx context.Context) ([]*permissionAppViews.PermissionView, error) {
	var items []models.Permission
	if err := q.db.WithContext(ctx).
		Where("deleted_at IS NULL").
		Order("category ASC, tag ASC").
		Find(&items).Error; err != nil {
		return nil, err
	}
	return slice.MapP(items, func(m *models.Permission) *permissionAppViews.PermissionView {
		view := mapper.PermissionModelToAppView(m)
		return &view
	}), nil
}

