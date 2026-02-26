package account_lockouts

import (
	"nfxidentity/modules/auth/domain/account_lockouts"
	"nfxidentity/modules/auth/infrastructure/repository/account_lockouts/check"
	"nfxidentity/modules/auth/infrastructure/repository/account_lockouts/create"
	"nfxidentity/modules/auth/infrastructure/repository/account_lockouts/delete"
	"nfxidentity/modules/auth/infrastructure/repository/account_lockouts/get"
	"nfxidentity/modules/auth/infrastructure/repository/account_lockouts/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 AccountLockout repository
func NewRepo(db *gorm.DB) *account_lockouts.Repo {
	return &account_lockouts.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
