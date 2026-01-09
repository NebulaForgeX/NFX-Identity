package create

import (
	"nfxid/modules/auth/domain/password_resets"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 password_resets.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) password_resets.Create {
	return &Handler{db: db}
}
