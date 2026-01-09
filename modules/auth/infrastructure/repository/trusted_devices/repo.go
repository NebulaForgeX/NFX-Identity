package trusted_devices

import (
	"nfxid/modules/auth/domain/trusted_devices"
	"nfxid/modules/auth/infrastructure/repository/trusted_devices/check"
	"nfxid/modules/auth/infrastructure/repository/trusted_devices/create"
	"nfxid/modules/auth/infrastructure/repository/trusted_devices/delete"
	"nfxid/modules/auth/infrastructure/repository/trusted_devices/get"
	"nfxid/modules/auth/infrastructure/repository/trusted_devices/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 TrustedDevice repository
func NewRepo(db *gorm.DB) *trusted_devices.Repo {
	return &trusted_devices.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
