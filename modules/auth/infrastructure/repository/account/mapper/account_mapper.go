package mapper

import (
	accountDomain "nfxidentity/modules/auth/domain/account"
	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/pkgs/ptrx"
)

func AccountDomainToModel(u *accountDomain.Account) *rdbmodels.Account {
	if u == nil {
		return nil
	}
	return &rdbmodels.Account{
		ID:             u.ID(),
		AccountStatus:  u.AccountStatus(),
		SignupPlatform: u.SignupPlatform(),
		CreatedAt:      u.CreatedAt(),
		UpdatedAt:      u.UpdatedAt(),
		DeletedAt:      ptrx.TimePtrToDeletedAt(u.DeletedAt()),
	}
}

func AccountModelToDomain(m *rdbmodels.Account) *accountDomain.Account {
	if m == nil {
		return nil
	}
	return accountDomain.NewAccountFromState(accountDomain.AccountState{
		ID:             m.ID,
		AccountStatus:  m.AccountStatus,
		SignupPlatform: m.SignupPlatform,
		CreatedAt:      m.CreatedAt,
		UpdatedAt:      m.UpdatedAt,
		DeletedAt:      ptrx.DeletedAtToTimePtr(m.DeletedAt),
	})
}

func AccountDomainToUpdates(u *accountDomain.Account) map[string]any {
	m := AccountDomainToModel(u)
	return map[string]any{
		rdbmodels.AccountCols.AccountStatus: m.AccountStatus,
		rdbmodels.AccountCols.UpdatedAt:     m.UpdatedAt,
		rdbmodels.AccountCols.DeletedAt:     m.DeletedAt,
	}
}
