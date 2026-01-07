package create

import (
	userRole "nfxid/modules/auth/domain/user_role"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 user_role.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) userRole.Create {
	return &Handler{db: db}
}
