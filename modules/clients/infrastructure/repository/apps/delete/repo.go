package delete

import (
	"nfxid/modules/clients/domain/apps"

	"gorm.io/gorm"
)

// Handler 处理删除操作，实现 apps.Delete 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Delete Handler
func NewHandler(db *gorm.DB) apps.Delete {
	return &Handler{db: db}
}
