package check

import (
	"nfxid/modules/auth/domain/profile"

	"gorm.io/gorm"
)

// Handler 处理检查操作，实现 profile.Check 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Check Handler
func NewHandler(db *gorm.DB) profile.Check {
	return &Handler{db: db}
}
