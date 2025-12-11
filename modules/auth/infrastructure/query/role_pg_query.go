package query

import (
	"context"
	"errors"
	roleAppQueries "nfxid/modules/auth/application/role/queries"
	roleDomainErrors "nfxid/modules/auth/domain/role/errors"
	roleDomainViews "nfxid/modules/auth/domain/role/views"
	"nfxid/modules/auth/infrastructure/query/mapper"
	"nfxid/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type rolePGQuery struct {
	db *gorm.DB
}

func NewRolePGQuery(db *gorm.DB) *rolePGQuery {
	return &rolePGQuery{db: db}
}

func (q *rolePGQuery) GetByID(ctx context.Context, id uuid.UUID) (roleDomainViews.RoleView, error) {
	var m models.Role
	if err := q.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return roleDomainViews.RoleView{}, roleDomainErrors.ErrRoleViewNotFound
		}
		return roleDomainViews.RoleView{}, err
	}
	return mapper.RoleModelToDomain(&m), nil
}

func (q *rolePGQuery) GetByName(ctx context.Context, name string) (roleDomainViews.RoleView, error) {
	var m models.Role
	if err := q.db.WithContext(ctx).Where("name = ?", name).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return roleDomainViews.RoleView{}, roleDomainErrors.ErrRoleViewNotFound
		}
		return roleDomainViews.RoleView{}, err
	}
	return mapper.RoleModelToDomain(&m), nil
}

func (q *rolePGQuery) GetList(ctx context.Context, listQuery roleAppQueries.RoleListQuery) ([]roleDomainViews.RoleView, int64, error) {
	var items []models.Role
	var total int64

	listQuery.Normalize()
	commonQuery := mapper.RoleListQueryToCommonQuery(listQuery)

	queryBuilder := q.db.WithContext(ctx).Model(&models.Role{})

	// Apply search
	if commonQuery.Search != "" {
		queryBuilder = queryBuilder.Where(
			"name ILIKE ? OR description ILIKE ?",
			"%"+commonQuery.Search+"%",
			"%"+commonQuery.Search+"%",
		)
	}

	// Apply filters
	if listQuery.IsSystem != nil {
		queryBuilder = queryBuilder.Where("is_system = ?", *listQuery.IsSystem)
	}

	// Count total
	if err := queryBuilder.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination
	if !commonQuery.All {
		if commonQuery.Offset > 0 {
			queryBuilder = queryBuilder.Offset(commonQuery.Offset)
		}
		if commonQuery.Limit > 0 {
			queryBuilder = queryBuilder.Limit(commonQuery.Limit)
		}
	}

	// Apply sorting
	if len(commonQuery.Sorts) > 0 {
		for _, sort := range commonQuery.Sorts {
			queryBuilder = queryBuilder.Order(sort.Field + " " + sort.Order)
		}
	} else {
		queryBuilder = queryBuilder.Order("created_at DESC")
	}

	// Execute query
	if err := queryBuilder.Find(&items).Error; err != nil {
		return nil, 0, err
	}

	// Convert to domain views
	result := make([]roleDomainViews.RoleView, len(items))
	for i, item := range items {
		result[i] = mapper.RoleModelToDomain(&item)
	}

	return result, total, nil
}

func (q *rolePGQuery) GetCount(ctx context.Context) (int64, error) {
	var count int64
	if err := q.db.WithContext(ctx).Model(&models.Role{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
