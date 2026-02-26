package tenant_roles

import (
	"nfxidentity/modules/access/domain/tenant_roles"
	"nfxidentity/modules/access/infrastructure/repository/tenant_roles/check"
	"nfxidentity/modules/access/infrastructure/repository/tenant_roles/create"
	"nfxidentity/modules/access/infrastructure/repository/tenant_roles/delete"
	"nfxidentity/modules/access/infrastructure/repository/tenant_roles/get"
	"nfxidentity/modules/access/infrastructure/repository/tenant_roles/update"

	"gorm.io/gorm"
)

func NewRepo(db *gorm.DB) *tenant_roles.Repo {
	return &tenant_roles.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
