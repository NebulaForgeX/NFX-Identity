package forger

import (
	"context"

	"nfxidentity/modules/auth/infrastructure/rdb/views"
	profileQuery "nfxidentity/modules/auth/query/profile"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) profileQuery.Forger { return &Handler{db: db} }

func toVO(r views.ListForgerProfileItem) profileQuery.ForgerItemVO {
	return profileQuery.ForgerItemVO{
		ProfileID: r.ProfileID, AccountID: r.AccountID, ForgerRoles: r.ForgerRoles,
		DisplayName: r.DisplayName, ProfileLanguage: r.ProfileLanguage, City: r.City, Country: r.Country,
		Website: r.Website, Timezone: r.Timezone, AvatarImageID: r.AvatarImageID, CreatedAt: r.CreatedAt,
	}
}

func (h *Handler) ByAccountID(ctx context.Context, accountID uuid.UUID) ([]profileQuery.ForgerItemVO, error) {
	var rows []views.ListForgerProfileItem
	if err := h.db.WithContext(ctx).Table(views.ListForgerProfileItem{}.TableName()).
		Where("account_id = ?", accountID).Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]profileQuery.ForgerItemVO, 0, len(rows))
	for _, r := range rows {
		out = append(out, toVO(r))
	}
	return out, nil
}

func (h *Handler) Search(ctx context.Context, q string, limit, offset int) ([]profileQuery.ForgerItemVO, int64, error) {
	tx := h.db.WithContext(ctx).Table(views.ListForgerProfileItem{}.TableName())
	if q != "" {
		like := "%" + q + "%"
		tx = tx.Where("display_name ILIKE ? OR city ILIKE ? OR country ILIKE ?", like, like, like)
	}
	var total int64
	if err := tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []views.ListForgerProfileItem
	if err := tx.Order("created_at DESC").Limit(limit).Offset(offset).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	out := make([]profileQuery.ForgerItemVO, 0, len(rows))
	for _, r := range rows {
		out = append(out, toVO(r))
	}
	return out, total, nil
}
