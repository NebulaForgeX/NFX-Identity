package refresh_tokens

import (
	"nfxidentity/modules/auth/domain/refresh_tokens"
	"nfxidentity/modules/auth/infrastructure/repository/refresh_tokens/check"
	"nfxidentity/modules/auth/infrastructure/repository/refresh_tokens/create"
	"nfxidentity/modules/auth/infrastructure/repository/refresh_tokens/delete"
	"nfxidentity/modules/auth/infrastructure/repository/refresh_tokens/get"
	"nfxidentity/modules/auth/infrastructure/repository/refresh_tokens/update"

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
