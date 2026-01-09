package update

import (
	"nfxid/modules/auth/domain/login_attempts"

	"gorm.io/gorm"
)

// Handler 处理更新操作，实现 login_attempts.Update 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Update Handler
func NewHandler(db *gorm.DB) login_attempts.Update {
	return &Handler{db: db}
}
