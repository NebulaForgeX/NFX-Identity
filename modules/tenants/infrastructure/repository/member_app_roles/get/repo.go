package get

import (
	"nfxid/modules/tenants/domain/member_app_roles"

	"gorm.io/gorm"
)

// Handler 处理获取数据操作，实现 member_app_roles.Get 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Get Handler
func NewHandler(db *gorm.DB) member_app_roles.Get {
	return &Handler{db: db}
}
