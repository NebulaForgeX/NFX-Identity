package ip_allowlist

import (
	"nfxidentity/modules/clients/domain/ip_allowlist"
	"nfxidentity/modules/clients/infrastructure/repository/ip_allowlist/check"
	"nfxidentity/modules/clients/infrastructure/repository/ip_allowlist/create"
	"nfxidentity/modules/clients/infrastructure/repository/ip_allowlist/delete"
	"nfxidentity/modules/clients/infrastructure/repository/ip_allowlist/get"
	"nfxidentity/modules/clients/infrastructure/repository/ip_allowlist/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 IPAllowlist repository
func NewRepo(db *gorm.DB) *ip_allowlist.Repo {
	return &ip_allowlist.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
