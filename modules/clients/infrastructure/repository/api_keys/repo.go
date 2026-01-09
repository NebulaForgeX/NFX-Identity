package api_keys

import (
	"nfxid/modules/clients/domain/api_keys"
	"nfxid/modules/clients/infrastructure/repository/api_keys/check"
	"nfxid/modules/clients/infrastructure/repository/api_keys/create"
	"nfxid/modules/clients/infrastructure/repository/api_keys/delete"
	"nfxid/modules/clients/infrastructure/repository/api_keys/get"
	"nfxid/modules/clients/infrastructure/repository/api_keys/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 APIKey repository
func NewRepo(db *gorm.DB) *api_keys.Repo {
	return &api_keys.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
