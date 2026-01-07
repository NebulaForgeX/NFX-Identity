package single

import (
	badgeDomain "nfxid/modules/auth/domain/badge"

	"gorm.io/gorm"
)

// Handler 处理单个查询操作，实现 badgeDomain.Single 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Single Handler
func NewHandler(db *gorm.DB) badgeDomain.Single {
	return &Handler{db: db}
}
