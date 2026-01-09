package check

import (
	"nfxid/modules/auth/domain/account_lockouts"

	"gorm.io/gorm"
)

// Handler 处理检查操作，实现 account_lockouts.Check 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Check Handler
func NewHandler(db *gorm.DB) account_lockouts.Check {
	return &Handler{db: db}
}
