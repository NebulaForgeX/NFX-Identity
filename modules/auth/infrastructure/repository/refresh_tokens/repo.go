package refresh_tokens

import (
	"nfxid/modules/auth/domain/refresh_tokens"
	"nfxid/modules/auth/infrastructure/repository/refresh_tokens/check"
	"nfxid/modules/auth/infrastructure/repository/refresh_tokens/create"
	"nfxid/modules/auth/infrastructure/repository/refresh_tokens/delete"
	"nfxid/modules/auth/infrastructure/repository/refresh_tokens/get"
	"nfxid/modules/auth/infrastructure/repository/refresh_tokens/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 RefreshToken repository
func NewRepo(db *gorm.DB) *refresh_tokens.Repo {
	return &refresh_tokens.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
