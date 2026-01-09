package check

import (
	"nfxid/modules/directory/domain/user_badges"

	"gorm.io/gorm"
)

// Handler 处理检查操作，实现 user_badges.Check 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Check Handler
func NewHandler(db *gorm.DB) user_badges.Check {
	return &Handler{db: db}
}
