package mapper

import (
	"nfxidentity/enums"
	"nfxidentity/modules/auth/domain/account"
	"nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/pkgs/utils/timex"
)

func ToModel(a *account.Account) *models.Account {
	st := a.State()
	return &models.Account{
		ID:             st.ID,
		AccountStatus:  enums.AuthAccountStatus(st.AccountStatus),
		SignupPlatform: enums.AuthSignupPlatform(st.SignupPlatform),
		CreatedAt:      st.CreatedAt,
		UpdatedAt:      st.UpdatedAt,
		DeletedAt:      timex.TimeToGormDeletedAt(st.DeletedAt),
	}
}

func ToDomain(m *models.Account) *account.Account {
	return account.NewFromState(account.AccountState{
		ID:             m.ID,
		AccountStatus:  string(m.AccountStatus),
		SignupPlatform: string(m.SignupPlatform),
		CreatedAt:      m.CreatedAt,
		UpdatedAt:      m.UpdatedAt,
		DeletedAt:      timex.GormDeletedAtToTime(m.DeletedAt),
	})
}
