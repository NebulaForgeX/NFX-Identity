package get

import (
	permissionDomain "nfxid/modules/permission/domain/permission"

	"gorm.io/gorm"
)

// Handler 处理获取数据操作，实现 permissionDomain.Get 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Get Handler
func NewHandler(db *gorm.DB) permissionDomain.Get {
	return &Handler{db: db}
}
