package update

import (
	"nfxid/modules/tenants/domain/member_roles"

	"gorm.io/gorm"
)

// Handler 处理更新操作，实现 member_roles.Update 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Update Handler
func NewHandler(db *gorm.DB) member_roles.Update {
	return &Handler{db: db}
}
