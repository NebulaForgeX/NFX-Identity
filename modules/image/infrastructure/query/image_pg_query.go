package query

import (
	"context"
	"encoding/json"
	"errors"
	imageAppQueries "nfxid/modules/image/application/image/queries"
	imageDomainErrors "nfxid/modules/image/domain/image/errors"
	imageDomainViews "nfxid/modules/image/domain/image/views"
	"nfxid/modules/image/infrastructure/rdb/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type imagePGQuery struct {
	db *gorm.DB
}

func NewImagePGQuery(db *gorm.DB) *imagePGQuery {
	return &imagePGQuery{db: db}
}

func (q *imagePGQuery) GetByID(ctx context.Context, id uuid.UUID) (imageDomainViews.ImageView, error) {
	var m models.Image
	if err := q.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return imageDomainViews.ImageView{}, imageDomainErrors.ErrImageNotFound
		}
		return imageDomainViews.ImageView{}, err
	}
	return imageModelToDomainView(&m), nil
}

func (q *imagePGQuery) GetList(ctx context.Context, listQuery imageAppQueries.ImageListQuery) ([]imageDomainViews.ImageView, int64, error) {
	var items []models.Image
	var total int64

	listQuery.Normalize()
	offset := (listQuery.Page - 1) * listQuery.PageSize

	queryBuilder := q.db.WithContext(ctx).Model(&models.Image{})

	// Apply filters
	if listQuery.TypeID != nil {
		queryBuilder = queryBuilder.Where("type_id = ?", *listQuery.TypeID)
	}
	if listQuery.UserID != nil {
		queryBuilder = queryBuilder.Where("user_id = ?", *listQuery.UserID)
	}
	if listQuery.SourceDomain != nil {
		queryBuilder = queryBuilder.Where("source_domain = ?", *listQuery.SourceDomain)
	}
	if listQuery.IsPublic != nil {
		queryBuilder = queryBuilder.Where("is_public = ?", *listQuery.IsPublic)
	}
	if listQuery.Search != "" {
		queryBuilder = queryBuilder.Where(
			"filename ILIKE ? OR original_filename ILIKE ?",
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

	views := make([]imageDomainViews.ImageView, len(items))
	for i, item := range items {
		views[i] = imageModelToDomainView(&item)
	}

	return views, total, nil
}

func imageModelToDomainView(m *models.Image) imageDomainViews.ImageView {
	var metadata map[string]interface{}
	if m.Metadata != nil {
		rawJSON, _ := m.Metadata.MarshalJSON()
		if len(rawJSON) > 0 {
			_ = json.Unmarshal(rawJSON, &metadata)
		}
	}

	return imageDomainViews.ImageView{
		ID:               m.ID,
		TypeID:           m.TypeID,
		UserID:           m.UserID,
		SourceDomain:     m.SourceDomain,
		Filename:         m.Filename,
		OriginalFilename: m.OriginalFilename,
		MimeType:         m.MimeType,
		Size:             m.Size,
		Width:            m.Width,
		Height:           m.Height,
		StoragePath:      m.StoragePath,
		URL:              m.URL,
		IsPublic:         m.IsPublic,
		Metadata:         metadata,
		CreatedAt:        m.CreatedAt,
		UpdatedAt:        m.UpdatedAt,
	}
}
