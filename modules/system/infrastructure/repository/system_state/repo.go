package system_state

import (
	"nfxidentity/modules/system/domain/system_state"
	"nfxidentity/modules/system/infrastructure/repository/system_state/check"
	"nfxidentity/modules/system/infrastructure/repository/system_state/create"
	"nfxidentity/modules/system/infrastructure/repository/system_state/delete"
	"nfxidentity/modules/system/infrastructure/repository/system_state/get"
	"nfxidentity/modules/system/infrastructure/repository/system_state/update"

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
