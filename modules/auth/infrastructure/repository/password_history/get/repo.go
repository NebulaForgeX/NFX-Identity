package get

import (
	"nfxid/modules/auth/domain/password_history"

	"gorm.io/gorm"
)

// Handler 处理获取数据操作，实现 password_history.Get 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Get Handler
func NewHandler(db *gorm.DB) password_history.Get {
	return &Handler{db: db}
}
