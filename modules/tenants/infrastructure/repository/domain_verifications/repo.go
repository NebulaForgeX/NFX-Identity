package domain_verifications

import (
	"nfxid/modules/tenants/domain/domain_verifications"
	"nfxid/modules/tenants/infrastructure/repository/domain_verifications/check"
	"nfxid/modules/tenants/infrastructure/repository/domain_verifications/create"
	"nfxid/modules/tenants/infrastructure/repository/domain_verifications/delete"
	"nfxid/modules/tenants/infrastructure/repository/domain_verifications/get"
	"nfxid/modules/tenants/infrastructure/repository/domain_verifications/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 DomainVerification repository
func NewRepo(db *gorm.DB) *domain_verifications.Repo {
	return &domain_verifications.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
