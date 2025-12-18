package single

import (
	roleDomain "nfxid/modules/auth/domain/role"

	"gorm.io/gorm"
)

// Handler 处理单个查询操作，实现 roleDomain.Single 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Single Handler
func NewHandler(db *gorm.DB) roleDomain.Single {
	return &Handler{db: db}
}
