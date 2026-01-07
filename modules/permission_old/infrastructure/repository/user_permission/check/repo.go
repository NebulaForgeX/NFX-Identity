package check

import (
	userPermissionDomain "nfxid/modules/permission/domain/user_permission"

	"gorm.io/gorm"
)

// Handler 处理检查操作，实现 userPermissionDomain.Check 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Check Handler
func NewHandler(db *gorm.DB) userPermissionDomain.Check {
	return &Handler{db: db}
}
