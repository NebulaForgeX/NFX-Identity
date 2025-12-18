package role

import (
	roleDomain "nfxid/modules/auth/domain/role"
	"nfxid/modules/auth/infrastructure/query/role/list"
	"nfxid/modules/auth/infrastructure/query/role/single"

	"gorm.io/gorm"
)

// NewHandler 创建新的 Role Query Handler
func NewHandler(db *gorm.DB) *roleDomain.Query {
	return &roleDomain.Query{
		Single: single.NewHandler(db),
		List:   list.NewHandler(db),
	}
}
