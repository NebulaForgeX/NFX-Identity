package update

import (
	"nfxid/modules/access/domain/role_permissions"

	"gorm.io/gorm"
)

// Handler 处理更新操作，实现 role_permissions.Update 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Update Handler
func NewHandler(db *gorm.DB) role_permissions.Update {
	return &Handler{db: db}
}
