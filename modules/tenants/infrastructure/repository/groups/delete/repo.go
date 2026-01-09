package delete

import (
	"nfxid/modules/tenants/domain/groups"

	"gorm.io/gorm"
)

// Handler 处理删除操作，实现 groups.Delete 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Delete Handler
func NewHandler(db *gorm.DB) groups.Delete {
	return &Handler{db: db}
}
