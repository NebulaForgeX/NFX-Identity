package count

import (
	userRoleDomain "nfxid/modules/auth/domain/user_role"

	"gorm.io/gorm"
)

// Handler 处理计数操作，实现 userRoleDomain.Count 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Count Handler
func NewHandler(db *gorm.DB) userRoleDomain.Count {
	return &Handler{db: db}
}
