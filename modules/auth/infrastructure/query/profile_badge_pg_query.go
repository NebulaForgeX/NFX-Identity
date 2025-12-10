package query

import (
	"context"
	"errors"
	profileBadgeAppQueries "nebulaid/modules/auth/application/profile_badge/queries"
	profileBadgeDomainErrors "nebulaid/modules/auth/domain/profile_badge/errors"
	profileBadgeDomainViews "nebulaid/modules/auth/domain/profile_badge/views"
	userDomainViews "nebulaid/modules/auth/domain/user/views"
	"nebulaid/modules/auth/infrastructure/query/mapper"
	"nebulaid/modules/auth/infrastructure/rdb/models"
	"nebulaid/modules/auth/infrastructure/rdb/views"
	"nebulaid/pkgs/utils/slice"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type profileBadgePGQuery struct {
	db *gorm.DB
}

func NewProfileBadgePGQuery(db *gorm.DB) *profileBadgePGQuery {
	return &profileBadgePGQuery{db: db}
}

func (q *profileBadgePGQuery) GetByID(ctx context.Context, id uuid.UUID) (profileBadgeDomainViews.ProfileBadgeView, error) {
	var m models.ProfileBadge
	if err := q.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return profileBadgeDomainViews.ProfileBadgeView{}, profileBadgeDomainErrors.ErrProfileBadgeViewNotFound
		}
		return profileBadgeDomainViews.ProfileBadgeView{}, err
	}
	return mapper.ProfileBadgeModelToDomain(&m), nil
}

func (q *profileBadgePGQuery) GetByProfileID(ctx context.Context, profileID uuid.UUID) ([]profileBadgeDomainViews.ProfileBadgeView, error) {
	var items []models.ProfileBadge
	if err := q.db.WithContext(ctx).
		Where("profile_id = ?", profileID).
		Order("earned_at DESC").
		Find(&items).Error; err != nil {
		return nil, err
	}
	return slice.MapP(items, mapper.ProfileBadgeModelToDomain), nil
}

func (q *profileBadgePGQuery) GetByBadgeID(ctx context.Context, badgeID uuid.UUID) ([]profileBadgeDomainViews.ProfileBadgeView, error) {
	var items []models.ProfileBadge
	if err := q.db.WithContext(ctx).
		Where("badge_id = ?", badgeID).
		Order("earned_at DESC").
		Find(&items).Error; err != nil {
		return nil, err
	}
	return slice.MapP(items, mapper.ProfileBadgeModelToDomain), nil
}

func (q *profileBadgePGQuery) GetUserBadges(ctx context.Context, userID uuid.UUID) ([]userDomainViews.UserBadgesView, error) {
	var items []views.UserBadgesView
	if err := q.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Find(&items).Error; err != nil {
		return nil, err
	}
	return slice.MapP(items, mapper.UserBadgesViewToDomain), nil
}

func (q *profileBadgePGQuery) GetList(ctx context.Context, listQuery profileBadgeAppQueries.ProfileBadgeListQuery) ([]profileBadgeDomainViews.ProfileBadgeView, int64, error) {
	var items []models.ProfileBadge
	var total int64

	listQuery.Normalize()
	commonQuery := mapper.ProfileBadgeListQueryToCommonQuery(listQuery)

	queryBuilder := q.db.WithContext(ctx).Model(&models.ProfileBadge{})

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
		queryBuilder = queryBuilder.Order("earned_at DESC")
	}

	// Execute query
	if err := queryBuilder.Find(&items).Error; err != nil {
		return nil, 0, err
	}

	return slice.MapP(items, mapper.ProfileBadgeModelToDomain), total, nil
}

func (q *profileBadgePGQuery) GetCount(ctx context.Context) (int64, error) {
	var count int64
	if err := q.db.WithContext(ctx).Model(&models.ProfileBadge{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
