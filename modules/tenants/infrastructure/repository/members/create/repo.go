package create

import (
	"nfxid/modules/tenants/domain/members"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 members.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) members.Create {
	return &Handler{db: db}
}
