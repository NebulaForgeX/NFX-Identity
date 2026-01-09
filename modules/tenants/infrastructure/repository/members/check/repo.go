package check

import (
	"nfxid/modules/tenants/domain/members"

	"gorm.io/gorm"
)

// Handler 处理检查操作，实现 members.Check 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Check Handler
func NewHandler(db *gorm.DB) members.Check {
	return &Handler{db: db}
}
