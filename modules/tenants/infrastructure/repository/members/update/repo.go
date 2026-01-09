package update

import (
	"nfxid/modules/tenants/domain/members"

	"gorm.io/gorm"
)

// Handler 处理更新操作，实现 members.Update 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Update Handler
func NewHandler(db *gorm.DB) members.Update {
	return &Handler{db: db}
}
