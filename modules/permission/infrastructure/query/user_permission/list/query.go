package list

import (
	userPermissionDomain "nfxid/modules/permission/domain/user_permission"

	"gorm.io/gorm"
)

// Handler 处理列表查询操作，实现 userPermissionDomain.List 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 List Handler
func NewHandler(db *gorm.DB) userPermissionDomain.List {
	return &Handler{db: db}
}
