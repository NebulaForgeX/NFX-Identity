package single

import (
	userRoleDomain "nfxid/modules/auth/domain/user_role"

	"gorm.io/gorm"
)

// Handler 处理单个查询操作，实现 userRoleDomain.Single 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Single Handler
func NewHandler(db *gorm.DB) userRoleDomain.Single {
	return &Handler{db: db}
}
