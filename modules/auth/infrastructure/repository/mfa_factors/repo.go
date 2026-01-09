package mfa_factors

import (
	"nfxid/modules/auth/domain/mfa_factors"
	"nfxid/modules/auth/infrastructure/repository/mfa_factors/check"
	"nfxid/modules/auth/infrastructure/repository/mfa_factors/create"
	"nfxid/modules/auth/infrastructure/repository/mfa_factors/delete"
	"nfxid/modules/auth/infrastructure/repository/mfa_factors/get"
	"nfxid/modules/auth/infrastructure/repository/mfa_factors/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 MFAFactor repository
func NewRepo(db *gorm.DB) *mfa_factors.Repo {
	return &mfa_factors.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
