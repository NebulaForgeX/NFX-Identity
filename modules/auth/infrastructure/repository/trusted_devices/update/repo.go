package update

import (
	"nfxid/modules/auth/domain/trusted_devices"

	"gorm.io/gorm"
)

// Handler 处理更新操作，实现 trusted_devices.Update 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Update Handler
func NewHandler(db *gorm.DB) trusted_devices.Update {
	return &Handler{db: db}
}
