package application_roles

import (
	dom "nfxidentity/modules/access/domain/application_roles"
	"nfxidentity/modules/access/infrastructure/repository/application_roles/check"
	"nfxidentity/modules/access/infrastructure/repository/application_roles/create"
	"nfxidentity/modules/access/infrastructure/repository/application_roles/delete"
	"nfxidentity/modules/access/infrastructure/repository/application_roles/get"
	"nfxidentity/modules/access/infrastructure/repository/application_roles/update"

	"gorm.io/gorm"
)

func NewRepo(db *gorm.DB) *dom.Repo {
	return &dom.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
