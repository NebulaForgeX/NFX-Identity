package client_credentials

import (
	"nfxid/modules/clients/domain/client_credentials"
	"nfxid/modules/clients/infrastructure/repository/client_credentials/check"
	"nfxid/modules/clients/infrastructure/repository/client_credentials/create"
	"nfxid/modules/clients/infrastructure/repository/client_credentials/delete"
	"nfxid/modules/clients/infrastructure/repository/client_credentials/get"
	"nfxid/modules/clients/infrastructure/repository/client_credentials/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 ClientCredential repository
func NewRepo(db *gorm.DB) *client_credentials.Repo {
	return &client_credentials.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
