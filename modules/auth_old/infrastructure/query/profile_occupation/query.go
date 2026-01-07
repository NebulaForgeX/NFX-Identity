package profile_occupation

import (
	occupationDomain "nfxid/modules/auth/domain/profile_occupation"
	"nfxid/modules/auth/infrastructure/query/profile_occupation/count"
	"nfxid/modules/auth/infrastructure/query/profile_occupation/list"
	"nfxid/modules/auth/infrastructure/query/profile_occupation/single"

	"gorm.io/gorm"
)

// NewHandler 创建新的 Occupation Query Handler
func NewHandler(db *gorm.DB) *occupationDomain.Query {
	return &occupationDomain.Query{
		Single: single.NewHandler(db),
		List:   list.NewHandler(db),
		Count:  count.NewHandler(db),
	}
}
