package create

import (
	"nfxid/modules/clients/domain/ip_allowlist"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 ip_allowlist.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) ip_allowlist.Create {
	return &Handler{db: db}
}
