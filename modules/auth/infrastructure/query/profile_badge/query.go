package profile_badge

import (
	profileBadgeDomain "nfxid/modules/auth/domain/profile_badge"
	"nfxid/modules/auth/infrastructure/query/profile_badge/count"
	"nfxid/modules/auth/infrastructure/query/profile_badge/list"
	"nfxid/modules/auth/infrastructure/query/profile_badge/single"

	"gorm.io/gorm"
)

// NewHandler 创建新的 ProfileBadge Query Handler
func NewHandler(db *gorm.DB) *profileBadgeDomain.Query {
	return &profileBadgeDomain.Query{
		Single: single.NewHandler(db),
		List:   list.NewHandler(db),
		Count:  count.NewHandler(db),
	}
}
