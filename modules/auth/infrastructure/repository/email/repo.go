package email

import (
	emailDomain "nfxidentity/modules/auth/domain/email"
	"nfxidentity/modules/auth/infrastructure/repository/email/check"
	"nfxidentity/modules/auth/infrastructure/repository/email/create"
	"nfxidentity/modules/auth/infrastructure/repository/email/delete"
	"nfxidentity/modules/auth/infrastructure/repository/email/get"
	"nfxidentity/modules/auth/infrastructure/repository/email/update"

	"gorm.io/gorm"
)

func NewRepo(db *gorm.DB) *emailDomain.Repo {
	return &emailDomain.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
