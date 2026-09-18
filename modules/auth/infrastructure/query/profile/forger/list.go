package forger

import (
	"context"

	"nfxidentity/modules/auth/infrastructure/rdb/views"
	profileQuery "nfxidentity/modules/auth/query/profile"
	"nfxidentity/pkgs/utils/ptr"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) profileQuery.Forger { return &Handler{db: db} }

func toVO(r views.Listforgerprofileitem) profileQuery.ForgerItemVO {
	var roles pq.StringArray
	if r.ForgerRoles != nil {
		_ = roles.Scan(*r.ForgerRoles)
	}
	lang := ""
	if r.ProfileLanguage != nil {
		lang = string(*r.ProfileLanguage)
	}
	return profileQuery.ForgerItemVO{
		ProfileID: ptr.Deref(r.ProfileID), AccountID: ptr.Deref(r.AccountID), ForgerRoles: roles,
		DisplayName: r.DisplayName, ProfileLanguage: lang, City: r.City, Country: r.Country,
		Website: r.Website, Timezone: r.Timezone, AvatarImageID: r.AvatarImageID, CreatedAt: ptr.Deref(r.CreatedAt),
	}
}

func (h *Handler) ByAccountID(ctx context.Context, accountID uuid.UUID) ([]profileQuery.ForgerItemVO, error) {
	var rows []views.Listforgerprofileitem
	if err := h.db.WithContext(ctx).Table(views.Listforgerprofileitem{}.TableName()).
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
	tx := h.db.WithContext(ctx).Table(views.Listforgerprofileitem{}.TableName())
	if q != "" {
		like := "%" + q + "%"
		tx = tx.Where("display_name ILIKE ? OR city ILIKE ? OR country ILIKE ?", like, like, like)
	}
	var total int64
	if err := tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []views.Listforgerprofileitem
	if err := tx.Order("created_at DESC").Limit(limit).Offset(offset).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	out := make([]profileQuery.ForgerItemVO, 0, len(rows))
	for _, r := range rows {
		out = append(out, toVO(r))
	}
	return out, total, nil
}
