package rate_limits

import (
	"nfxidentity/modules/clients/domain/rate_limits"
	"nfxidentity/modules/clients/infrastructure/repository/rate_limits/check"
	"nfxidentity/modules/clients/infrastructure/repository/rate_limits/create"
	"nfxidentity/modules/clients/infrastructure/repository/rate_limits/delete"
	"nfxidentity/modules/clients/infrastructure/repository/rate_limits/get"
	"nfxidentity/modules/clients/infrastructure/repository/rate_limits/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 RateLimit repository
func NewRepo(db *gorm.DB) *rate_limits.Repo {
	return &rate_limits.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
