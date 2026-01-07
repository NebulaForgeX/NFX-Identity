package delete

import (
	"nfxid/modules/auth/domain/profile"

	"gorm.io/gorm"
)

// Handler 处理删除操作，实现 profile.Delete 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Delete Handler
func NewHandler(db *gorm.DB) profile.Delete {
	return &Handler{db: db}
}
