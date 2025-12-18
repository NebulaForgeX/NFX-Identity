package list

import (
	userDomain "nfxid/modules/auth/domain/user"

	"gorm.io/gorm"
)

// Handler 处理列表查询操作，实现 userDomain.List 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 List Handler
func NewHandler(db *gorm.DB) userDomain.List {
	return &Handler{db: db}
}
