package check

import (
	"nfxid/modules/auth/domain/sessions"

	"gorm.io/gorm"
)

// Handler 处理检查操作，实现 sessions.Check 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Check Handler
func NewHandler(db *gorm.DB) sessions.Check {
	return &Handler{db: db}
}
