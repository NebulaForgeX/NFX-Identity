package update

import (
	"nfxid/modules/auth/domain/password_history"

	"gorm.io/gorm"
)

// Handler 处理更新操作，实现 password_history.Update 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Update Handler
func NewHandler(db *gorm.DB) password_history.Update {
	return &Handler{db: db}
}
