package count

import (
	profileBadgeDomain "nfxid/modules/auth/domain/profile_badge"

	"gorm.io/gorm"
)

// Handler 处理计数操作，实现 profileBadgeDomain.Count 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Count Handler
func NewHandler(db *gorm.DB) profileBadgeDomain.Count {
	return &Handler{db: db}
}
