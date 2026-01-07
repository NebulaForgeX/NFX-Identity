package list

import (
	occupationDomain "nfxid/modules/auth/domain/profile_occupation"

	"gorm.io/gorm"
)

// Handler 处理列表查询操作，实现 occupationDomain.List 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 List Handler
func NewHandler(db *gorm.DB) occupationDomain.List {
	return &Handler{db: db}
}
