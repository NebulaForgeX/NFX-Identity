package create

import (
	"nfxid/modules/auth/domain/mfa_factors"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 mfa_factors.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) mfa_factors.Create {
	return &Handler{db: db}
}
