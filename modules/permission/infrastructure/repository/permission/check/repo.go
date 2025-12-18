package check

import (
	permissionDomain "nfxid/modules/permission/domain/permission"

	"gorm.io/gorm"
)

// Handler 处理检查操作，实现 permissionDomain.Check 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Check Handler
func NewHandler(db *gorm.DB) permissionDomain.Check {
	return &Handler{db: db}
}
