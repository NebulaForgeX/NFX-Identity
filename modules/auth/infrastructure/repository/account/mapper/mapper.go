package mapper

import (
	"nfxidentity/modules/auth/domain/account"
	"nfxidentity/modules/auth/infrastructure/rdb/models"
)

func ToModel(a *account.Account) *models.Account {
	st := a.State()
	return &models.Account{
		ID: st.ID, AccountStatus: st.AccountStatus, SignupPlatform: st.SignupPlatform,
		CreatedAt: st.CreatedAt, UpdatedAt: st.UpdatedAt, DeletedAt: st.DeletedAt,
	}
}

func ToDomain(m *models.Account) *account.Account {
	return account.NewFromState(account.AccountState{
		ID: m.ID, AccountStatus: m.AccountStatus, SignupPlatform: m.SignupPlatform,
		CreatedAt: m.CreatedAt, UpdatedAt: m.UpdatedAt, DeletedAt: m.DeletedAt,
	})
}
