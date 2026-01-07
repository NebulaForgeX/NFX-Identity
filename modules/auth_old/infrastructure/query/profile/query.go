package profile

import (
	profileDomain "nfxid/modules/auth/domain/profile"
	"nfxid/modules/auth/infrastructure/query/profile/count"
	"nfxid/modules/auth/infrastructure/query/profile/list"
	"nfxid/modules/auth/infrastructure/query/profile/single"

	"gorm.io/gorm"
)

// NewHandler 创建新的 Profile Query Handler
func NewHandler(db *gorm.DB) *profileDomain.Query {
	return &profileDomain.Query{
		Single: single.NewHandler(db),
		List:   list.NewHandler(db),
		Count:  count.NewHandler(db),
	}
}
