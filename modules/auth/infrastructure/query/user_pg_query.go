package query

import (
	"context"
	"errors"
	userAppQueries "nfxid/modules/auth/application/user/queries"
	userDomainErrors "nfxid/modules/auth/domain/user/errors"
	userDomainViews "nfxid/modules/auth/domain/user/views"
	"nfxid/modules/auth/infrastructure/query/mapper"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/rdb/views"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userPGQuery struct {
	db *gorm.DB
}

func NewUserPGQuery(db *gorm.DB) *userPGQuery {
	return &userPGQuery{db: db}
}

func (q *userPGQuery) GetByID(ctx context.Context, id uuid.UUID) (userDomainViews.UserView, error) {
	var v views.UserWithRoleView
	if err := q.db.WithContext(ctx).Where("user_id = ?", id).First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return userDomainViews.UserView{}, userDomainErrors.ErrUserViewNotFound
		}
		return userDomainViews.UserView{}, err
	}
	var u models.User
	if err := q.db.WithContext(ctx).Where("id = ?", id).First(&u).Error; err != nil {
		return userDomainViews.UserView{}, err
	}
	return mapper.UserViewToDomain(&v, &u), nil
}

func (q *userPGQuery) GetByUsername(ctx context.Context, username string) (userDomainViews.UserView, error) {
	var v views.UserWithRoleView
	if err := q.db.WithContext(ctx).Where("username = ?", username).First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return userDomainViews.UserView{}, userDomainErrors.ErrUserViewNotFound
		}
		return userDomainViews.UserView{}, err
	}
	var u models.User
	if err := q.db.WithContext(ctx).Where("username = ?", username).First(&u).Error; err != nil {
		return userDomainViews.UserView{}, err
	}
	return mapper.UserViewToDomain(&v, &u), nil
}

func (q *userPGQuery) GetByEmail(ctx context.Context, email string) (userDomainViews.UserView, error) {
	var v views.UserWithRoleView
	if err := q.db.WithContext(ctx).Where("email = ?", email).First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return userDomainViews.UserView{}, userDomainErrors.ErrUserViewNotFound
		}
		return userDomainViews.UserView{}, err
	}
	var u models.User
	if err := q.db.WithContext(ctx).Where("email = ?", email).First(&u).Error; err != nil {
		return userDomainViews.UserView{}, err
	}
	return mapper.UserViewToDomain(&v, &u), nil
}

func (q *userPGQuery) GetList(ctx context.Context, listQuery userAppQueries.UserListQuery) ([]userDomainViews.UserView, int64, error) {
	var items []views.UserWithRoleView
	var total int64

	listQuery.Normalize()
	commonQuery := mapper.UserListQueryToCommonQuery(listQuery)

	queryBuilder := q.db.WithContext(ctx).Model(&views.UserWithRoleView{})

	// Apply search
	if commonQuery.Search != "" {
		queryBuilder = queryBuilder.Where(
			"username ILIKE ? OR email ILIKE ? OR phone ILIKE ?",
			"%"+commonQuery.Search+"%",
			"%"+commonQuery.Search+"%",
			"%"+commonQuery.Search+"%",
		)
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
		queryBuilder = queryBuilder.Order("user_created_at DESC")
	}

	// Execute query
	if err := queryBuilder.Find(&items).Error; err != nil {
		return nil, 0, err
	}

	// Convert to domain views
	result := make([]userDomainViews.UserView, len(items))
	userIDs := make([]uuid.UUID, len(items))
	for i, item := range items {
		userIDs[i] = item.UserID
	}

	// Fetch all users in batch
	var users []models.User
	if err := q.db.WithContext(ctx).Where("id IN ?", userIDs).Find(&users).Error; err != nil {
		return nil, 0, err
	}
	userMap := make(map[uuid.UUID]*models.User)
	for i := range users {
		userMap[users[i].ID] = &users[i]
	}

	for i, item := range items {
		u := userMap[item.UserID]
		result[i] = mapper.UserViewToDomain(&item, u)
	}

	return result, total, nil
}

func (q *userPGQuery) GetCount(ctx context.Context) (int64, error) {
	var count int64
	if err := q.db.WithContext(ctx).Model(&models.User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
