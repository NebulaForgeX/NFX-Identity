package delete

import (
	"nfxid/modules/auth/domain/mfa_factors"

	"gorm.io/gorm"
)

// Handler 处理删除操作，实现 mfa_factors.Delete 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Delete Handler
func NewHandler(db *gorm.DB) mfa_factors.Delete {
	return &Handler{db: db}
}
