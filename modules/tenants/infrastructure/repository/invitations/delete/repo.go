package delete

import (
	"nfxid/modules/tenants/domain/invitations"

	"gorm.io/gorm"
)

// Handler 处理删除操作，实现 invitations.Delete 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Delete Handler
func NewHandler(db *gorm.DB) invitations.Delete {
	return &Handler{db: db}
}
