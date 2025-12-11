package query

import (
	"context"
	"errors"
	imageTypeAppQueries "nfxid/modules/image/application/image_type/queries"
	imageTypeDomainErrors "nfxid/modules/image/domain/image_type/errors"
	imageTypeDomainViews "nfxid/modules/image/domain/image_type/views"
	"nfxid/modules/image/infrastructure/rdb/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type imageTypePGQuery struct {
	db *gorm.DB
}

func NewImageTypePGQuery(db *gorm.DB) *imageTypePGQuery {
	return &imageTypePGQuery{db: db}
}

func (q *imageTypePGQuery) GetByID(ctx context.Context, id uuid.UUID) (imageTypeDomainViews.ImageTypeView, error) {
	var m models.ImageType
	if err := q.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return imageTypeDomainViews.ImageTypeView{}, imageTypeDomainErrors.ErrImageTypeNotFound
		}
		return imageTypeDomainViews.ImageTypeView{}, err
	}
	return imageTypeModelToDomainView(&m), nil
}

func (q *imageTypePGQuery) GetByKey(ctx context.Context, key string) (imageTypeDomainViews.ImageTypeView, error) {
	var m models.ImageType
	if err := q.db.WithContext(ctx).Where("key = ?", key).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return imageTypeDomainViews.ImageTypeView{}, imageTypeDomainErrors.ErrImageTypeNotFound
		}
		return imageTypeDomainViews.ImageTypeView{}, err
	}
	return imageTypeModelToDomainView(&m), nil
}

func (q *imageTypePGQuery) GetList(ctx context.Context, listQuery imageTypeAppQueries.ImageTypeListQuery) ([]imageTypeDomainViews.ImageTypeView, int64, error) {
	var items []models.ImageType
	var total int64

	listQuery.Normalize()
	offset := (listQuery.Page - 1) * listQuery.PageSize

	queryBuilder := q.db.WithContext(ctx).Model(&models.ImageType{})

	// Apply filters
	if listQuery.IsSystem != nil {
		queryBuilder = queryBuilder.Where("is_system = ?", *listQuery.IsSystem)
	}
	if listQuery.Search != "" {
		queryBuilder = queryBuilder.Where(
			"key ILIKE ? OR description ILIKE ?",
			"%"+listQuery.Search+"%",
			"%"+listQuery.Search+"%",
		)
	}

	// Count total
	if err := queryBuilder.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply sorting
	orderBy := listQuery.OrderBy
	if orderBy == "" {
		orderBy = "created_at"
	}
	order := listQuery.Order
	if order == "" {
		order = "desc"
	}
	queryBuilder = queryBuilder.Order(orderBy + " " + order)

	// Apply pagination
	if err := queryBuilder.Offset(offset).Limit(listQuery.PageSize).Find(&items).Error; err != nil {
		return nil, 0, err
	}

	views := make([]imageTypeDomainViews.ImageTypeView, len(items))
	for i, item := range items {
		views[i] = imageTypeModelToDomainView(&item)
	}

	return views, total, nil
}

func imageTypeModelToDomainView(m *models.ImageType) imageTypeDomainViews.ImageTypeView {
	isSystem := false
	if m.IsSystem != nil {
		isSystem = *m.IsSystem
	}

	return imageTypeDomainViews.ImageTypeView{
		ID:          m.ID,
		Key:         m.Key,
		Description: m.Description,
		MaxWidth:    m.MaxWidth,
		MaxHeight:   m.MaxHeight,
		AspectRatio: m.AspectRatio,
		IsSystem:    isSystem,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}
