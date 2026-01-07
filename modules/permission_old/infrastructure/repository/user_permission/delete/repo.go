package delete

import (
	userPermissionDomain "nfxid/modules/permission/domain/user_permission"

	"gorm.io/gorm"
)

// Handler 处理删除操作，实现 userPermissionDomain.Delete 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Delete Handler
func NewHandler(db *gorm.DB) userPermissionDomain.Delete {
	return &Handler{db: db}
}
