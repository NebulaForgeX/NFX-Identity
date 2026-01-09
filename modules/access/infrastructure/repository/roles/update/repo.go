package update

import (
	"nfxid/modules/access/domain/roles"

	"gorm.io/gorm"
)

// Handler 处理更新操作，实现 roles.Update 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Update Handler
func NewHandler(db *gorm.DB) roles.Update {
	return &Handler{db: db}
}
