package mapper

import (
	identityDomain "nfxidentity/modules/auth/domain/identity"
	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/pkgs/ptrx"
)

func IdentityDomainToModel(i *identityDomain.Identity) *rdbmodels.Identity {
	if i == nil {
		return nil
	}
	return &rdbmodels.Identity{
		ID:               i.ID(),
		AccountID:        i.AccountID(),
		IdentityProvider: i.IdentityProvider(),
		ProviderSubject:  i.ProviderSubject(),
		PasswordHash:     i.PasswordHash(),
		Metadata:         i.Metadata(),
		LastLoginAt:      i.LastLoginAt(),
		CreatedAt:        i.CreatedAt(),
		UpdatedAt:        i.UpdatedAt(),
		DeletedAt:        ptrx.TimePtrToDeletedAt(i.DeletedAt()),
	}
}

func IdentityModelToDomain(m *rdbmodels.Identity) *identityDomain.Identity {
	if m == nil {
		return nil
	}
	return identityDomain.NewIdentityFromState(identityDomain.IdentityState{
		ID:               m.ID,
		AccountID:        m.AccountID,
		IdentityProvider: m.IdentityProvider,
		ProviderSubject:  m.ProviderSubject,
		PasswordHash:     m.PasswordHash,
		Metadata:         m.Metadata,
		LastLoginAt:      m.LastLoginAt,
		CreatedAt:        m.CreatedAt,
		UpdatedAt:        m.UpdatedAt,
		DeletedAt:        ptrx.DeletedAtToTimePtr(m.DeletedAt),
	})
}

func IdentityDomainToUpdates(i *identityDomain.Identity) map[string]any {
	m := IdentityDomainToModel(i)
	return map[string]any{
		rdbmodels.IdentityCols.AccountID:        m.AccountID,
		rdbmodels.IdentityCols.IdentityProvider: m.IdentityProvider,
		rdbmodels.IdentityCols.ProviderSubject:  m.ProviderSubject,
		rdbmodels.IdentityCols.PasswordHash:     m.PasswordHash,
		rdbmodels.IdentityCols.Metadata:         m.Metadata,
		rdbmodels.IdentityCols.LastLoginAt:      m.LastLoginAt,
		rdbmodels.IdentityCols.UpdatedAt:        m.UpdatedAt,
		rdbmodels.IdentityCols.DeletedAt:        m.DeletedAt,
	}
}
