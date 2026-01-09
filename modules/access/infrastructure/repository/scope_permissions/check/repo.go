package check

import (
	"nfxid/modules/access/domain/scope_permissions"

	"gorm.io/gorm"
)

// Handler 处理检查操作，实现 scope_permissions.Check 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Check Handler
func NewHandler(db *gorm.DB) scope_permissions.Check {
	return &Handler{db: db}
}
