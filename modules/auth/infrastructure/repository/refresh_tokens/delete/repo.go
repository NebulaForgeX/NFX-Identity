package delete

import (
	"nfxid/modules/auth/domain/refresh_tokens"

	"gorm.io/gorm"
)

// Handler 处理删除操作，实现 refresh_tokens.Delete 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Delete Handler
func NewHandler(db *gorm.DB) refresh_tokens.Delete {
	return &Handler{db: db}
}
