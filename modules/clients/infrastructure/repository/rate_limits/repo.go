package rate_limits

import (
	"nfxid/modules/clients/domain/rate_limits"
	"nfxid/modules/clients/infrastructure/repository/rate_limits/check"
	"nfxid/modules/clients/infrastructure/repository/rate_limits/create"
	"nfxid/modules/clients/infrastructure/repository/rate_limits/delete"
	"nfxid/modules/clients/infrastructure/repository/rate_limits/get"
	"nfxid/modules/clients/infrastructure/repository/rate_limits/update"

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
