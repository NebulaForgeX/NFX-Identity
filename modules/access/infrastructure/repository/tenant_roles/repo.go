package tenant_roles

import (
	"nfxid/modules/access/domain/tenant_roles"
	"nfxid/modules/access/infrastructure/repository/tenant_roles/check"
	"nfxid/modules/access/infrastructure/repository/tenant_roles/create"
	"nfxid/modules/access/infrastructure/repository/tenant_roles/delete"
	"nfxid/modules/access/infrastructure/repository/tenant_roles/get"
	"nfxid/modules/access/infrastructure/repository/tenant_roles/update"
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
