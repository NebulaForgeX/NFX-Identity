package check

import (
	"nfxid/modules/auth/domain/password_resets"

	"gorm.io/gorm"
)

// Handler 处理检查操作，实现 password_resets.Check 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Check Handler
func NewHandler(db *gorm.DB) password_resets.Check {
	return &Handler{db: db}
}
