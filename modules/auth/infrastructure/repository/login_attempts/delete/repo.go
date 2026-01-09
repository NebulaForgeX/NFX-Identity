package delete

import (
	"nfxid/modules/auth/domain/login_attempts"

	"gorm.io/gorm"
)

// Handler 处理删除操作，实现 login_attempts.Delete 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Delete Handler
func NewHandler(db *gorm.DB) login_attempts.Delete {
	return &Handler{db: db}
}
