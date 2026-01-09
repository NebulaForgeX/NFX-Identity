package check

import (
	"nfxid/modules/auth/domain/refresh_tokens"

	"gorm.io/gorm"
)

// Handler 处理检查操作，实现 refresh_tokens.Check 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Check Handler
func NewHandler(db *gorm.DB) refresh_tokens.Check {
	return &Handler{db: db}
}
