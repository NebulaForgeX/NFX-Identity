package check

import (
	"nfxid/modules/auth/domain/role"

	"gorm.io/gorm"
)

// Handler 处理检查操作，实现 role.Check 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Check Handler
func NewHandler(db *gorm.DB) role.Check {
	return &Handler{db: db}
}
