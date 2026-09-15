package list

import (
	"context"

	emailQuery "nfxidentity/modules/auth/query/email"
	"nfxidentity/modules/auth/infrastructure/rdb/views"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) emailQuery.List { return &Handler{db: db} }

func (h *Handler) ByAccountID(ctx context.Context, accountID uuid.UUID) ([]emailQuery.EmailItemVO, error) {
	var rows []views.ListEmailItem
	if err := h.db.WithContext(ctx).
		Table(views.ListEmailItem{}.TableName()).
		Where("account_id = ?", accountID).
		Order("is_primary DESC, created_at ASC").
		Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]emailQuery.EmailItemVO, 0, len(rows))
	for _, r := range rows {
		out = append(out, emailQuery.EmailItemVO{
			ID: r.ID, AccountID: r.AccountID, Email: r.Email, IsPrimary: r.IsPrimary,
			VerifiedAt: r.VerifiedAt, CreatedAt: r.CreatedAt, UpdatedAt: r.UpdatedAt,
		})
	}
	return out, nil
}
