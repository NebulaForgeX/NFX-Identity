package list

import (
	profileBadgeDomain "nfxid/modules/auth/domain/profile_badge"

	"gorm.io/gorm"
)

// Handler 处理列表查询操作，实现 profileBadgeDomain.List 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 List Handler
func NewHandler(db *gorm.DB) profileBadgeDomain.List {
	return &Handler{db: db}
}
