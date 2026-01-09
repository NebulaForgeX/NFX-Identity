package create

import (
	"nfxid/modules/tenants/domain/groups"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 groups.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) groups.Create {
	return &Handler{db: db}
}
