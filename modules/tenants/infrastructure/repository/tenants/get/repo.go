package get

import (
	"nfxid/modules/tenants/domain/tenants"

	"gorm.io/gorm"
)

// Handler 处理获取数据操作，实现 tenants.Get 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Get Handler
func NewHandler(db *gorm.DB) tenants.Get {
	return &Handler{db: db}
}
