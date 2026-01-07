package permission

import (
	permissionDomain "nfxid/modules/permission/domain/permission"
	"nfxid/modules/permission/infrastructure/query/permission/list"
	"nfxid/modules/permission/infrastructure/query/permission/single"

	"gorm.io/gorm"
)

// NewHandler 创建新的 Permission Query Handler
func NewHandler(db *gorm.DB) *permissionDomain.Query {
	return &permissionDomain.Query{
		Single: single.NewHandler(db),
		List:   list.NewHandler(db),
	}
}
