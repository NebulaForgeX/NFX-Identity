package badge

import (
	badgeDomain "nfxid/modules/auth/domain/badge"
	"nfxid/modules/auth/infrastructure/query/badge/list"
	"nfxid/modules/auth/infrastructure/query/badge/single"

	"gorm.io/gorm"
)

// NewHandler 创建新的 Badge Query Handler
func NewHandler(db *gorm.DB) *badgeDomain.Query {
	return &badgeDomain.Query{
		Single: single.NewHandler(db),
		List:   list.NewHandler(db),
	}
}
