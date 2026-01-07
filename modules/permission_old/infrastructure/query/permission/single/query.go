package single

import (
	permissionDomain "nfxid/modules/permission/domain/permission"

	"gorm.io/gorm"
)

// Handler 处理单个查询操作，实现 permissionDomain.Single 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Single Handler
func NewHandler(db *gorm.DB) permissionDomain.Single {
	return &Handler{db: db}
}
