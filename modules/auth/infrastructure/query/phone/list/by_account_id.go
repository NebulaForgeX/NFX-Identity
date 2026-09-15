package list

import (
	"context"

	"nfxidentity/modules/auth/infrastructure/rdb/views"
	phoneQuery "nfxidentity/modules/auth/query/phone"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) phoneQuery.List { return &Handler{db: db} }

func (h *Handler) ByAccountID(ctx context.Context, accountID uuid.UUID) ([]phoneQuery.PhoneItemVO, error) {
	var rows []views.ListPhoneItem
	if err := h.db.WithContext(ctx).
		Table(views.ListPhoneItem{}.TableName()).
		Where("account_id = ?", accountID).
		Order("is_primary DESC, created_at ASC").
		Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]phoneQuery.PhoneItemVO, 0, len(rows))
	for _, r := range rows {
		out = append(out, phoneQuery.PhoneItemVO{
			ID: r.ID, AccountID: r.AccountID, Phone: r.Phone, IsPrimary: r.IsPrimary,
			VerifiedAt: r.VerifiedAt, CreatedAt: r.CreatedAt, UpdatedAt: r.UpdatedAt,
		})
	}
	return out, nil
}
