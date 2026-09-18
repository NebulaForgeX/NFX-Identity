package mapper

import (
	"nfxidentity/modules/auth/domain/identity"
	"nfxidentity/modules/auth/infrastructure/rdb/models"
)

func ToModel(i *identity.Identity) *models.Identity {
	st := i.State()
	return &models.Identity{ID: st.ID, AccountID: st.AccountID, IdentityProvider: st.IdentityProvider, ProviderSubject: st.ProviderSubject, PasswordHash: st.PasswordHash, LastLoginAt: st.LastLoginAt, CreatedAt: st.CreatedAt, UpdatedAt: st.UpdatedAt, DeletedAt: st.DeletedAt}
}
func ToDomain(m *models.Identity) *identity.Identity {
	return identity.NewFromState(identity.IdentityState{ID: m.ID, AccountID: m.AccountID, IdentityProvider: m.IdentityProvider, ProviderSubject: m.ProviderSubject, PasswordHash: m.PasswordHash, LastLoginAt: m.LastLoginAt, CreatedAt: m.CreatedAt, UpdatedAt: m.UpdatedAt, DeletedAt: m.DeletedAt})
}
