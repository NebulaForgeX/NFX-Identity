package get

import (
	"nfxid/modules/tenants/domain/member_groups"

	"gorm.io/gorm"
)

// Handler 处理获取数据操作，实现 member_groups.Get 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Get Handler
func NewHandler(db *gorm.DB) member_groups.Get {
	return &Handler{db: db}
}
