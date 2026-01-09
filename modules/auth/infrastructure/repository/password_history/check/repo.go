package check

import (
	"nfxid/modules/auth/domain/password_history"

	"gorm.io/gorm"
)

// Handler 处理检查操作，实现 password_history.Check 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Check Handler
func NewHandler(db *gorm.DB) password_history.Check {
	return &Handler{db: db}
}
