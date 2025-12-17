package user_role

import (
	userRole "nfxid/modules/auth/domain/user_role"
	"nfxid/modules/auth/infrastructure/repository/user_role/check"
	"nfxid/modules/auth/infrastructure/repository/user_role/create"
	"nfxid/modules/auth/infrastructure/repository/user_role/delete"
	"nfxid/modules/auth/infrastructure/repository/user_role/get"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 UserRole repository
func NewRepo(db *gorm.DB) *userRole.Repo {
	return &userRole.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
