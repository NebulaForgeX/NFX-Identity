package user_role

import (
	userRoleDomain "nfxid/modules/auth/domain/user_role"

	"gorm.io/gorm"
)

// Handler 处理查询操作，实现 user_role.Query 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 UserRole Query Handler
func NewHandler(db *gorm.DB) userRoleDomain.Query {
	return &Handler{db: db}
}
