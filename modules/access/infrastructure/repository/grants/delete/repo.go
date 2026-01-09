package delete

import (
	"nfxid/modules/access/domain/grants"

	"gorm.io/gorm"
)

// Handler 处理删除操作，实现 grants.Delete 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Delete Handler
func NewHandler(db *gorm.DB) grants.Delete {
	return &Handler{db: db}
}
