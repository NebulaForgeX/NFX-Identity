package occupation

import (
	occupation "nfxid/modules/auth/domain/profile_occupation"
	"nfxid/modules/auth/infrastructure/repository/profile_occupation/check"
	"nfxid/modules/auth/infrastructure/repository/profile_occupation/create"
	"nfxid/modules/auth/infrastructure/repository/profile_occupation/delete"
	"nfxid/modules/auth/infrastructure/repository/profile_occupation/get"
	"nfxid/modules/auth/infrastructure/repository/profile_occupation/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 Occupation repository
func NewRepo(db *gorm.DB) *occupation.Repo {
	return &occupation.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
