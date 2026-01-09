package account_lockouts

import (
	"nfxid/modules/auth/domain/account_lockouts"
	"nfxid/modules/auth/infrastructure/repository/account_lockouts/check"
	"nfxid/modules/auth/infrastructure/repository/account_lockouts/create"
	"nfxid/modules/auth/infrastructure/repository/account_lockouts/delete"
	"nfxid/modules/auth/infrastructure/repository/account_lockouts/get"
	"nfxid/modules/auth/infrastructure/repository/account_lockouts/update"

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
