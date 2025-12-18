package user

import (
	"nfxid/modules/auth/domain/user"
	"nfxid/modules/auth/infrastructure/query/user/count"
	"nfxid/modules/auth/infrastructure/query/user/list"
	"nfxid/modules/auth/infrastructure/query/user/single"

	"gorm.io/gorm"
)

// NewHandler 创建新的 User Query Handler
func NewHandler(db *gorm.DB) *user.Query {
	return &user.Query{
		Single: single.NewHandler(db),
		List:   list.NewHandler(db),
		Count:  count.NewHandler(db),
	}
}
