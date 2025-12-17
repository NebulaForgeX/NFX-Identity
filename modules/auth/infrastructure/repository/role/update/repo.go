package update

import (
	"nfxid/modules/auth/domain/role"

	"gorm.io/gorm"
)

// Handler 处理更新操作，实现 role.Update 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Update Handler
func NewHandler(db *gorm.DB) role.Update {
	return &Handler{db: db}
}
