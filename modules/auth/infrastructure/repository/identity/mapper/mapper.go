package mapper

import (
	"nfxidentity/enums"
	"nfxidentity/modules/auth/domain/identity"
	"nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/pkgs/utils/timex"
)

func ToModel(i *identity.Identity) *models.Identity {
	st := i.State()
	return &models.Identity{
		ID:               st.ID,
		AccountID:        st.AccountID,
		IdentityProvider: enums.AuthIdentityProvider(st.IdentityProvider),
		ProviderSubject:  st.ProviderSubject,
		PasswordHash:     st.PasswordHash,
		LastLoginAt:      st.LastLoginAt,
		CreatedAt:        st.CreatedAt,
		UpdatedAt:        st.UpdatedAt,
		DeletedAt:        timex.TimeToGormDeletedAt(st.DeletedAt),
	}
}

func ToDomain(m *models.Identity) *identity.Identity {
	return identity.NewFromState(identity.IdentityState{
		ID:               m.ID,
		AccountID:        m.AccountID,
		IdentityProvider: string(m.IdentityProvider),
		ProviderSubject:  m.ProviderSubject,
		PasswordHash:     m.PasswordHash,
		LastLoginAt:      m.LastLoginAt,
		CreatedAt:        m.CreatedAt,
		UpdatedAt:        m.UpdatedAt,
		DeletedAt:        timex.GormDeletedAtToTime(m.DeletedAt),
	})
}
