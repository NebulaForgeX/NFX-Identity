package ip_allowlist

import (
	"nfxid/modules/clients/domain/ip_allowlist"
	"nfxid/modules/clients/infrastructure/repository/ip_allowlist/check"
	"nfxid/modules/clients/infrastructure/repository/ip_allowlist/create"
	"nfxid/modules/clients/infrastructure/repository/ip_allowlist/delete"
	"nfxid/modules/clients/infrastructure/repository/ip_allowlist/get"
	"nfxid/modules/clients/infrastructure/repository/ip_allowlist/update"

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
