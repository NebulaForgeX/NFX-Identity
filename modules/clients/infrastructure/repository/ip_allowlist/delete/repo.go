package delete

import (
	"nfxid/modules/clients/domain/ip_allowlist"

	"gorm.io/gorm"
)

// Handler 处理删除操作，实现 ip_allowlist.Delete 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Delete Handler
func NewHandler(db *gorm.DB) ip_allowlist.Delete {
	return &Handler{db: db}
}
