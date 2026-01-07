package create

import (
	userPermissionDomain "nfxid/modules/permission/domain/user_permission"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 userPermissionDomain.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) userPermissionDomain.Create {
	return &Handler{db: db}
}
