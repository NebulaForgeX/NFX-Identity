package query

import (
	"context"
	"errors"
	userRoleAppQueries "nfxid/modules/auth/application/user_role/queries"
	userRoleDomainErrors "nfxid/modules/auth/domain/user_role/errors"
	userRoleDomainViews "nfxid/modules/auth/domain/user_role/views"
	"nfxid/modules/auth/infrastructure/query/mapper"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/pkgs/utils/slice"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRolePGQuery struct {
	db *gorm.DB
}

func NewUserRolePGQuery(db *gorm.DB) *userRolePGQuery {
	return &userRolePGQuery{db: db}
}

func (q *userRolePGQuery) GetByID(ctx context.Context, id uuid.UUID) (userRoleDomainViews.UserRoleView, error) {
	var m models.UserRole
	if err := q.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return userRoleDomainViews.UserRoleView{}, userRoleDomainErrors.ErrUserRoleViewNotFound
		}
		return userRoleDomainViews.UserRoleView{}, err
	}
	return mapper.UserRoleModelToDomain(&m), nil
}

func (q *userRolePGQuery) GetByUserID(ctx context.Context, userID uuid.UUID) ([]userRoleDomainViews.UserRoleView, error) {
	var items []models.UserRole
	if err := q.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&items).Error; err != nil {
		return nil, err
	}
	return slice.MapP(items, mapper.UserRoleModelToDomain), nil
}

func (q *userRolePGQuery) GetByRoleID(ctx context.Context, roleID uuid.UUID) ([]userRoleDomainViews.UserRoleView, error) {
	var items []models.UserRole
	if err := q.db.WithContext(ctx).
		Where("role_id = ?", roleID).
		Order("created_at DESC").
		Find(&items).Error; err != nil {
		return nil, err
	}
	return slice.MapP(items, mapper.UserRoleModelToDomain), nil
}

func (q *userRolePGQuery) GetList(ctx context.Context, listQuery userRoleAppQueries.UserRoleListQuery) ([]userRoleDomainViews.UserRoleView, int64, error) {
	var items []models.UserRole
	var total int64

	listQuery.Normalize()
	commonQuery := mapper.UserRoleListQueryToCommonQuery(listQuery)

	queryBuilder := q.db.WithContext(ctx).Model(&models.UserRole{})

	// Apply filters
	if len(listQuery.UserIDs) > 0 {
		queryBuilder = queryBuilder.Where("user_id IN ?", listQuery.UserIDs)
	}
	if len(listQuery.RoleIDs) > 0 {
		queryBuilder = queryBuilder.Where("role_id IN ?", listQuery.RoleIDs)
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

	return slice.MapP(items, mapper.UserRoleModelToDomain), total, nil
}

func (q *userRolePGQuery) GetCount(ctx context.Context) (int64, error) {
	var count int64
	if err := q.db.WithContext(ctx).Model(&models.UserRole{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

