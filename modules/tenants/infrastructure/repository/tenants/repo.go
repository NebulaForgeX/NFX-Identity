package tenants

import (
	"nfxidentity/modules/tenants/domain/tenants"
	"nfxidentity/modules/tenants/infrastructure/repository/tenants/check"
	"nfxidentity/modules/tenants/infrastructure/repository/tenants/create"
	"nfxidentity/modules/tenants/infrastructure/repository/tenants/delete"
	"nfxidentity/modules/tenants/infrastructure/repository/tenants/get"
	"nfxidentity/modules/tenants/infrastructure/repository/tenants/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 Tenant repository
func NewRepo(db *gorm.DB) *tenants.Repo {
	return &tenants.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
