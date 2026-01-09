package delete

import (
	"nfxid/modules/access/domain/role_permissions"

	"gorm.io/gorm"
)

// Handler 处理删除操作，实现 role_permissions.Delete 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Delete Handler
func NewHandler(db *gorm.DB) role_permissions.Delete {
	return &Handler{db: db}
}
