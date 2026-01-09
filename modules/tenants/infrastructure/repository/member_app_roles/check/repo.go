package check

import (
	"nfxid/modules/tenants/domain/member_app_roles"

	"gorm.io/gorm"
)

// Handler 处理检查操作，实现 member_app_roles.Check 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Check Handler
func NewHandler(db *gorm.DB) member_app_roles.Check {
	return &Handler{db: db}
}
