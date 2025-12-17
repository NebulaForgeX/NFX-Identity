package badge

import (
	"nfxid/modules/auth/domain/badge"
	"nfxid/modules/auth/infrastructure/repository/badge/create"
	"nfxid/modules/auth/infrastructure/repository/badge/get"
	"nfxid/modules/auth/infrastructure/repository/badge/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 Badge repository
func NewRepo(db *gorm.DB) *badge.Repo {
	return &badge.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Update: update.NewHandler(db),
	}
}
