package mapper

import (
	"nfxidentity/modules/auth/domain/phone"
	"nfxidentity/modules/auth/infrastructure/rdb/models"
)

func ToModel(p *phone.Phone) *models.Phone {
	st := p.State()
	return &models.Phone{ID: st.ID, AccountID: st.AccountID, Phone: st.Number, IsPrimary: st.IsPrimary, VerifiedAt: st.VerifiedAt, CreatedAt: st.CreatedAt, UpdatedAt: st.UpdatedAt, DeletedAt: st.DeletedAt}
}
func ToDomain(m *models.Phone) *phone.Phone {
	return phone.NewFromState(phone.PhoneState{ID: m.ID, AccountID: m.AccountID, Number: m.Phone, IsPrimary: m.IsPrimary, VerifiedAt: m.VerifiedAt, CreatedAt: m.CreatedAt, UpdatedAt: m.UpdatedAt, DeletedAt: m.DeletedAt})
}
