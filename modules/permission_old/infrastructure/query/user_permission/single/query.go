package single

import (
	userPermissionDomain "nfxid/modules/permission/domain/user_permission"

	"gorm.io/gorm"
)

// Handler 处理单个查询操作，实现 userPermissionDomain.Single 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Single Handler
func NewHandler(db *gorm.DB) userPermissionDomain.Single {
	return &Handler{db: db}
}
