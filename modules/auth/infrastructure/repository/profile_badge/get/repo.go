package get

import (
	profileBadge "nfxid/modules/auth/domain/profile_badge"

	"gorm.io/gorm"
)

// Handler 处理获取数据操作，实现 profileBadge.Get 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Get Handler
func NewHandler(db *gorm.DB) profileBadge.Get {
	return &Handler{db: db}
}
