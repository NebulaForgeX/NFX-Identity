package check

import (
	"nfxid/modules/auth/domain/login_attempts"

	"gorm.io/gorm"
)

// Handler 处理检查操作，实现 login_attempts.Check 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Check Handler
func NewHandler(db *gorm.DB) login_attempts.Check {
	return &Handler{db: db}
}
