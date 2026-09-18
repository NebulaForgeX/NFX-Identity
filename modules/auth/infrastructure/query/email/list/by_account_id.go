package list

import (
	"context"

	"nfxidentity/modules/auth/infrastructure/rdb/views"
	emailQuery "nfxidentity/modules/auth/query/email"
	"nfxidentity/pkgs/utils/ptr"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) emailQuery.List { return &Handler{db: db} }

func (h *Handler) ByAccountID(ctx context.Context, accountID uuid.UUID) ([]emailQuery.EmailItemVO, error) {
	var rows []views.Listemailitem
	if err := h.db.WithContext(ctx).
		Table(views.Listemailitem{}.TableName()).
		Where("account_id = ?", accountID).
		Order("is_primary DESC, created_at ASC").
		Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]emailQuery.EmailItemVO, 0, len(rows))
	for _, r := range rows {
		out = append(out, emailQuery.EmailItemVO{
			ID: ptr.Deref(r.ID), AccountID: ptr.Deref(r.AccountID), Email: ptr.Deref(r.Email),
			IsPrimary: ptr.Deref(r.IsPrimary), VerifiedAt: r.VerifiedAt,
			CreatedAt: ptr.Deref(r.CreatedAt), UpdatedAt: ptr.Deref(r.UpdatedAt),
		})
	}
	return out, nil
}
