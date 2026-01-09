package tenants

import (
	"nfxid/modules/tenants/domain/tenants"
	"nfxid/modules/tenants/infrastructure/repository/tenants/check"
	"nfxid/modules/tenants/infrastructure/repository/tenants/create"
	"nfxid/modules/tenants/infrastructure/repository/tenants/delete"
	"nfxid/modules/tenants/infrastructure/repository/tenants/get"
	"nfxid/modules/tenants/infrastructure/repository/tenants/update"

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
