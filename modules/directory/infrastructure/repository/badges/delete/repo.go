package delete

import (
	"nfxid/modules/directory/domain/badges"

	"gorm.io/gorm"
)

// Handler 处理删除操作，实现 badges.Delete 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Delete Handler
func NewHandler(db *gorm.DB) badges.Delete {
	return &Handler{db: db}
}
