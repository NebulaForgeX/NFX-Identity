package check

import (
	userRole "nfxid/modules/auth/domain/user_role"

	"gorm.io/gorm"
)

// Handler 处理检查操作，实现 user_role.Check 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Check Handler
func NewHandler(db *gorm.DB) userRole.Check {
	return &Handler{db: db}
}
