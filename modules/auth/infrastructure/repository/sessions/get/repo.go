package get

import (
	"nfxid/modules/auth/domain/sessions"

	"gorm.io/gorm"
)

// Handler 处理获取数据操作，实现 sessions.Get 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Get Handler
func NewHandler(db *gorm.DB) sessions.Get {
	return &Handler{db: db}
}
