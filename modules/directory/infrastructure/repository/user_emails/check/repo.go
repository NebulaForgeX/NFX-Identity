package check

import (
	"nfxid/modules/directory/domain/user_emails"

	"gorm.io/gorm"
)

// Handler 处理检查操作，实现 user_emails.Check 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Check Handler
func NewHandler(db *gorm.DB) user_emails.Check {
	return &Handler{db: db}
}
