package get

import (
	"nfxid/modules/tenants/domain/groups"

	"gorm.io/gorm"
)

// Handler 处理获取数据操作，实现 groups.Get 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Get Handler
func NewHandler(db *gorm.DB) groups.Get {
	return &Handler{db: db}
}
