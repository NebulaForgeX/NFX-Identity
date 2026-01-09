package update

import (
	"nfxid/modules/tenants/domain/invitations"

	"gorm.io/gorm"
)

// Handler 处理更新操作，实现 invitations.Update 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Update Handler
func NewHandler(db *gorm.DB) invitations.Update {
	return &Handler{db: db}
}
