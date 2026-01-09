package create

import (
	"nfxid/modules/auth/domain/refresh_tokens"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 refresh_tokens.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) refresh_tokens.Create {
	return &Handler{db: db}
}
