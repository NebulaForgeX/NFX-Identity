package tenant_apps

import (
	"nfxidentity/modules/tenants/domain/tenant_apps"
	"nfxidentity/modules/tenants/infrastructure/repository/tenant_apps/check"
	"nfxidentity/modules/tenants/infrastructure/repository/tenant_apps/create"
	"nfxidentity/modules/tenants/infrastructure/repository/tenant_apps/delete"
	"nfxidentity/modules/tenants/infrastructure/repository/tenant_apps/get"
	"nfxidentity/modules/tenants/infrastructure/repository/tenant_apps/update"

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
