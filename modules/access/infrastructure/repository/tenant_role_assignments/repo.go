package tenant_role_assignments

import (
	dom "nfxid/modules/access/domain/tenant_role_assignments"
	"nfxid/modules/access/infrastructure/repository/tenant_role_assignments/check"
	"nfxid/modules/access/infrastructure/repository/tenant_role_assignments/create"
	"nfxid/modules/access/infrastructure/repository/tenant_role_assignments/delete"
	"nfxid/modules/access/infrastructure/repository/tenant_role_assignments/get"
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
