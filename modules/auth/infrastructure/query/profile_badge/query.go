package profile_badge

import (
	profileBadgeDomain "nfxid/modules/auth/domain/profile_badge"

	"gorm.io/gorm"
)

// Handler 处理查询操作，实现 profile_badge.Query 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 ProfileBadge Query Handler
func NewHandler(db *gorm.DB) profileBadgeDomain.Query {
	return &Handler{db: db}
}
