package create

import (
	"nfxid/modules/tenants/domain/member_app_roles"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 member_app_roles.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) member_app_roles.Create {
	return &Handler{db: db}
}
