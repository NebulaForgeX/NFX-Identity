package delete

import (
	userRole "nfxid/modules/auth/domain/user_role"

	"gorm.io/gorm"
)

// Handler 处理删除操作，实现 user_role.Delete 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Delete Handler
func NewHandler(db *gorm.DB) userRole.Delete {
	return &Handler{db: db}
}
