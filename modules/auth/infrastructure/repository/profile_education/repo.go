package education

import (
	education "nfxid/modules/auth/domain/profile_education"
	"nfxid/modules/auth/infrastructure/repository/profile_education/check"
	"nfxid/modules/auth/infrastructure/repository/profile_education/create"
	"nfxid/modules/auth/infrastructure/repository/profile_education/delete"
	"nfxid/modules/auth/infrastructure/repository/profile_education/get"
	"nfxid/modules/auth/infrastructure/repository/profile_education/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 Education repository
func NewRepo(db *gorm.DB) *education.Repo {
	return &education.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
