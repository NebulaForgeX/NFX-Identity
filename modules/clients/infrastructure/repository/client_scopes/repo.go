package client_scopes

import (
	"nfxid/modules/clients/domain/client_scopes"
	"nfxid/modules/clients/infrastructure/repository/client_scopes/check"
	"nfxid/modules/clients/infrastructure/repository/client_scopes/create"
	"nfxid/modules/clients/infrastructure/repository/client_scopes/delete"
	"nfxid/modules/clients/infrastructure/repository/client_scopes/get"
	"nfxid/modules/clients/infrastructure/repository/client_scopes/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 ClientScope repository
func NewRepo(db *gorm.DB) *client_scopes.Repo {
	return &client_scopes.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
