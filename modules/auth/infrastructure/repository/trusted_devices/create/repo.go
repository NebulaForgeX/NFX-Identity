package create

import (
	"nfxid/modules/auth/domain/trusted_devices"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 trusted_devices.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) trusted_devices.Create {
	return &Handler{db: db}
}
