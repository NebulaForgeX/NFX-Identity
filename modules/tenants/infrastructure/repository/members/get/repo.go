package get

import (
	"nfxid/modules/tenants/domain/members"

	"gorm.io/gorm"
)

// Handler 处理获取数据操作，实现 members.Get 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Get Handler
func NewHandler(db *gorm.DB) members.Get {
	return &Handler{db: db}
}
