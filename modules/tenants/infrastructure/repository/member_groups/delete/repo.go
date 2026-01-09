package delete

import (
	"nfxid/modules/tenants/domain/member_groups"

	"gorm.io/gorm"
)

// Handler 处理删除操作，实现 member_groups.Delete 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Delete Handler
func NewHandler(db *gorm.DB) member_groups.Delete {
	return &Handler{db: db}
}
