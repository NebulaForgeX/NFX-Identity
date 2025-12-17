package get

import (
	userRole "nfxid/modules/auth/domain/user_role"

	"gorm.io/gorm"
)

// Handler 处理获取数据操作，实现 user_role.Get 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Get Handler
func NewHandler(db *gorm.DB) userRole.Get {
	return &Handler{db: db}
}
