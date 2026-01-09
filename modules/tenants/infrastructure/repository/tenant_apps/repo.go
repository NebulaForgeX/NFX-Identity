package tenant_apps

import (
	"nfxid/modules/tenants/domain/tenant_apps"
	"nfxid/modules/tenants/infrastructure/repository/tenant_apps/check"
	"nfxid/modules/tenants/infrastructure/repository/tenant_apps/create"
	"nfxid/modules/tenants/infrastructure/repository/tenant_apps/delete"
	"nfxid/modules/tenants/infrastructure/repository/tenant_apps/get"
	"nfxid/modules/tenants/infrastructure/repository/tenant_apps/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 TenantApp repository
func NewRepo(db *gorm.DB) *tenant_apps.Repo {
	return &tenant_apps.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
