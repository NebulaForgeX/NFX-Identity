package profile_education

import (
	educationDomain "nfxid/modules/auth/domain/profile_education"
	"nfxid/modules/auth/infrastructure/query/profile_education/count"
	"nfxid/modules/auth/infrastructure/query/profile_education/list"
	"nfxid/modules/auth/infrastructure/query/profile_education/single"

	"gorm.io/gorm"
)

// NewHandler 创建新的 Education Query Handler
func NewHandler(db *gorm.DB) *educationDomain.Query {
	return &educationDomain.Query{
		Single: single.NewHandler(db),
		List:   list.NewHandler(db),
		Count:  count.NewHandler(db),
	}
}
