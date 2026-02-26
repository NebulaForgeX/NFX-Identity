package tenant_settings

import (
	"nfxidentity/modules/tenants/domain/tenant_settings"
	"nfxidentity/modules/tenants/infrastructure/repository/tenant_settings/check"
	"nfxidentity/modules/tenants/infrastructure/repository/tenant_settings/create"
	"nfxidentity/modules/tenants/infrastructure/repository/tenant_settings/delete"
	"nfxidentity/modules/tenants/infrastructure/repository/tenant_settings/get"
	"nfxidentity/modules/tenants/infrastructure/repository/tenant_settings/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 TenantSetting repository
func NewRepo(db *gorm.DB) *tenant_settings.Repo {
	return &tenant_settings.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
