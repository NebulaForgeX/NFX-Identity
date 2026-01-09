package system_state

import (
	"nfxid/modules/system/domain/system_state"
	"nfxid/modules/system/infrastructure/repository/system_state/check"
	"nfxid/modules/system/infrastructure/repository/system_state/create"
	"nfxid/modules/system/infrastructure/repository/system_state/delete"
	"nfxid/modules/system/infrastructure/repository/system_state/get"
	"nfxid/modules/system/infrastructure/repository/system_state/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 SystemState repository
func NewRepo(db *gorm.DB) *system_state.Repo {
	return &system_state.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
