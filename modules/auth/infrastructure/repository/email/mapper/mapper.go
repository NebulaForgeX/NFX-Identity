package mapper

import (
	"nfxidentity/modules/auth/domain/email"
	"nfxidentity/modules/auth/infrastructure/rdb/models"
)

func ToModel(e *email.Email) *models.Email {
	st := e.State()
	return &models.Email{ID: st.ID, AccountID: st.AccountID, Email: st.Address, IsPrimary: st.IsPrimary, VerifiedAt: st.VerifiedAt, CreatedAt: st.CreatedAt, UpdatedAt: st.UpdatedAt, DeletedAt: st.DeletedAt}
}
func ToDomain(m *models.Email) *email.Email {
	return email.NewFromState(email.EmailState{ID: m.ID, AccountID: m.AccountID, Address: m.Email, IsPrimary: m.IsPrimary, VerifiedAt: m.VerifiedAt, CreatedAt: m.CreatedAt, UpdatedAt: m.UpdatedAt, DeletedAt: m.DeletedAt})
}
