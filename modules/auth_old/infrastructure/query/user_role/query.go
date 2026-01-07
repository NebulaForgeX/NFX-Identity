package user_role

import (
	userRoleDomain "nfxid/modules/auth/domain/user_role"
	"nfxid/modules/auth/infrastructure/query/user_role/count"
	"nfxid/modules/auth/infrastructure/query/user_role/list"
	"nfxid/modules/auth/infrastructure/query/user_role/single"

	"gorm.io/gorm"
)

// NewHandler 创建新的 UserRole Query Handler
func NewHandler(db *gorm.DB) *userRoleDomain.Query {
	return &userRoleDomain.Query{
		Single: single.NewHandler(db),
		List:   list.NewHandler(db),
		Count:  count.NewHandler(db),
	}
}
