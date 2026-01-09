package create

import (
	"nfxid/modules/auth/domain/login_attempts"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 login_attempts.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) login_attempts.Create {
	return &Handler{db: db}
}
