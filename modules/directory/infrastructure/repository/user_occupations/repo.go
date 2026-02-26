package user_occupations

import (
	"nfxidentity/modules/directory/domain/user_occupations"
	"nfxidentity/modules/directory/infrastructure/repository/user_occupations/check"
	"nfxidentity/modules/directory/infrastructure/repository/user_occupations/create"
	"nfxidentity/modules/directory/infrastructure/repository/user_occupations/delete"
	"nfxidentity/modules/directory/infrastructure/repository/user_occupations/get"
	"nfxidentity/modules/directory/infrastructure/repository/user_occupations/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 UserOccupation repository
func NewRepo(db *gorm.DB) *user_occupations.Repo {
	return &user_occupations.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
