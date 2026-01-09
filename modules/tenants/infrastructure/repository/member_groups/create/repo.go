package create

import (
	"nfxid/modules/tenants/domain/member_groups"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 member_groups.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) member_groups.Create {
	return &Handler{db: db}
}
