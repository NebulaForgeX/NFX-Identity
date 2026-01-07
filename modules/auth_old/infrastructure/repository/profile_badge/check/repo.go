package check

import (
	profileBadge "nfxid/modules/auth/domain/profile_badge"

	"gorm.io/gorm"
)

// Handler 处理检查操作，实现 profileBadge.Check 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Check Handler
func NewHandler(db *gorm.DB) profileBadge.Check {
	return &Handler{db: db}
}
