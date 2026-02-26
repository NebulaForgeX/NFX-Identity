package mfa_factors

import (
	"nfxidentity/modules/auth/domain/mfa_factors"
	"nfxidentity/modules/auth/infrastructure/repository/mfa_factors/check"
	"nfxidentity/modules/auth/infrastructure/repository/mfa_factors/create"
	"nfxidentity/modules/auth/infrastructure/repository/mfa_factors/delete"
	"nfxidentity/modules/auth/infrastructure/repository/mfa_factors/get"
	"nfxidentity/modules/auth/infrastructure/repository/mfa_factors/update"

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
