package delete

import (
	"nfxid/modules/auth/domain/password_history"

	"gorm.io/gorm"
)

// Handler 处理删除操作，实现 password_history.Delete 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Delete Handler
func NewHandler(db *gorm.DB) password_history.Delete {
	return &Handler{db: db}
}
