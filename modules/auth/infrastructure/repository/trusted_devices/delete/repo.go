package delete

import (
	"nfxid/modules/auth/domain/trusted_devices"

	"gorm.io/gorm"
)

// Handler 处理删除操作，实现 trusted_devices.Delete 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Delete Handler
func NewHandler(db *gorm.DB) trusted_devices.Delete {
	return &Handler{db: db}
}
