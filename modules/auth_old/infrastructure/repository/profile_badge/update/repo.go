package update

import (
	profileBadge "nfxid/modules/auth/domain/profile_badge"

	"gorm.io/gorm"
)

// Handler 处理更新操作，实现 profileBadge.Update 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Update Handler
func NewHandler(db *gorm.DB) profileBadge.Update {
	return &Handler{db: db}
}
