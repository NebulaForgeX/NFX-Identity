package update

import (
	"nfxid/modules/auth/domain/refresh_tokens"

	"gorm.io/gorm"
)

// Handler 处理更新操作，实现 refresh_tokens.Update 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Update Handler
func NewHandler(db *gorm.DB) refresh_tokens.Update {
	return &Handler{db: db}
}
