package list

import (
	roleDomain "nfxid/modules/auth/domain/role"

	"gorm.io/gorm"
)

// Handler 处理列表查询操作，实现 roleDomain.List 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 List Handler
func NewHandler(db *gorm.DB) roleDomain.List {
	return &Handler{db: db}
}
