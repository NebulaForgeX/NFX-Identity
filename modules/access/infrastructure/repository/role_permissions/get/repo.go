package get

import (
	"nfxid/modules/access/domain/role_permissions"

	"gorm.io/gorm"
)

// Handler 处理获取数据操作，实现 role_permissions.Get 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Get Handler
func NewHandler(db *gorm.DB) role_permissions.Get {
	return &Handler{db: db}
}
