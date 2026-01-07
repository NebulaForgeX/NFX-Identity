package count

import (
	userDomain "nfxid/modules/auth/domain/user"

	"gorm.io/gorm"
)

// Handler 处理计数操作，实现 userDomain.Count 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Count Handler
func NewHandler(db *gorm.DB) userDomain.Count {
	return &Handler{db: db}
}
