package authorization_code

import (
	authorizationCodeDomain "nfxid/modules/permission/domain/authorization_code"
	"nfxid/modules/permission/infrastructure/repository/authorization_code/check"
	"nfxid/modules/permission/infrastructure/repository/authorization_code/create"
	"nfxid/modules/permission/infrastructure/repository/authorization_code/delete"
	"nfxid/modules/permission/infrastructure/repository/authorization_code/get"
	"nfxid/modules/permission/infrastructure/repository/authorization_code/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 AuthorizationCode repository
func NewRepo(db *gorm.DB) *authorizationCodeDomain.Repo {
	return &authorizationCodeDomain.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
