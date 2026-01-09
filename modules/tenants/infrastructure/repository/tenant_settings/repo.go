package tenant_settings

import (
	"nfxid/modules/tenants/domain/tenant_settings"
	"nfxid/modules/tenants/infrastructure/repository/tenant_settings/check"
	"nfxid/modules/tenants/infrastructure/repository/tenant_settings/create"
	"nfxid/modules/tenants/infrastructure/repository/tenant_settings/delete"
	"nfxid/modules/tenants/infrastructure/repository/tenant_settings/get"
	"nfxid/modules/tenants/infrastructure/repository/tenant_settings/update"

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
