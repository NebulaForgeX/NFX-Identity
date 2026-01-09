package update

import (
	"nfxid/modules/clients/domain/ip_allowlist"

	"gorm.io/gorm"
)

// Handler 处理更新操作，实现 ip_allowlist.Update 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Update Handler
func NewHandler(db *gorm.DB) ip_allowlist.Update {
	return &Handler{db: db}
}
