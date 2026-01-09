package delete

import (
	"nfxid/modules/tenants/domain/member_app_roles"

	"gorm.io/gorm"
)

// Handler 处理删除操作，实现 member_app_roles.Delete 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Delete Handler
func NewHandler(db *gorm.DB) member_app_roles.Delete {
	return &Handler{db: db}
}
