package application_role_assignments

import (
	dom "nfxidentity/modules/access/domain/application_role_assignments"
	"nfxidentity/modules/access/infrastructure/repository/application_role_assignments/check"
	"nfxidentity/modules/access/infrastructure/repository/application_role_assignments/create"
	"nfxidentity/modules/access/infrastructure/repository/application_role_assignments/delete"
	"nfxidentity/modules/access/infrastructure/repository/application_role_assignments/get"

	"gorm.io/gorm"
)

func NewRepo(db *gorm.DB) *dom.Repo {
	return &dom.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
