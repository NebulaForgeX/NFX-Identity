package delete

import (
	"nfxid/modules/auth/domain/role"

	"gorm.io/gorm"
)

// Handler 处理删除操作，实现 role.Delete 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Delete Handler
func NewHandler(db *gorm.DB) role.Delete {
	return &Handler{db: db}
}
