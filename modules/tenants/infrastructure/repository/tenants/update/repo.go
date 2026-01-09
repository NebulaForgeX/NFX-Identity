package update

import (
	"nfxid/modules/tenants/domain/tenants"

	"gorm.io/gorm"
)

// Handler 处理更新操作，实现 tenants.Update 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Update Handler
func NewHandler(db *gorm.DB) tenants.Update {
	return &Handler{db: db}
}
