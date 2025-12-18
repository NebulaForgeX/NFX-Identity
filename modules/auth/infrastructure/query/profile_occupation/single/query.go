package single

import (
	occupationDomain "nfxid/modules/auth/domain/profile_occupation"

	"gorm.io/gorm"
)

// Handler 处理单个查询操作，实现 occupationDomain.Single 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Single Handler
func NewHandler(db *gorm.DB) occupationDomain.Single {
	return &Handler{db: db}
}
