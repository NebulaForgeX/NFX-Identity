package authority

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

func NewHandler(db *gorm.DB) profileQuery.Authority { return &Handler{db: db} }

func toVO(r views.Listauthorityprofileitem) profileQuery.AuthorityItemVO {
	var roles pq.StringArray
	if r.AuthorityRoles != nil {
		_ = roles.Scan(*r.AuthorityRoles)
	}
	lang := ""
	if r.ProfileLanguage != nil {
		lang = string(*r.ProfileLanguage)
	}
	return profileQuery.AuthorityItemVO{
		ProfileID: ptr.Deref(r.ProfileID), AccountID: ptr.Deref(r.AccountID), AuthorityRoles: roles,
		DisplayName: r.DisplayName, ProfileLanguage: lang, City: r.City, Country: r.Country,
		AvatarImageID: r.AvatarImageID, CreatedAt: ptr.Deref(r.CreatedAt),
	}
}

func (h *Handler) ByAccountID(ctx context.Context, accountID uuid.UUID) ([]profileQuery.AuthorityItemVO, error) {
	var rows []views.Listauthorityprofileitem
	if err := h.db.WithContext(ctx).Table(views.Listauthorityprofileitem{}.TableName()).
		Where("account_id = ?", accountID).Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]profileQuery.AuthorityItemVO, 0, len(rows))
	for _, r := range rows {
		out = append(out, toVO(r))
	}
	return out, nil
}

func (h *Handler) Search(ctx context.Context, q string, limit, offset int) ([]profileQuery.AuthorityItemVO, int64, error) {
	tx := h.db.WithContext(ctx).Table(views.Listauthorityprofileitem{}.TableName())
	if q != "" {
		like := "%" + q + "%"
		tx = tx.Where("display_name ILIKE ? OR city ILIKE ? OR country ILIKE ?", like, like, like)
	}
	var total int64
	if err := tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []views.Listauthorityprofileitem
	if err := tx.Order("created_at DESC").Limit(limit).Offset(offset).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	out := make([]profileQuery.AuthorityItemVO, 0, len(rows))
	for _, r := range rows {
		out = append(out, toVO(r))
	}
	return out, total, nil
}
