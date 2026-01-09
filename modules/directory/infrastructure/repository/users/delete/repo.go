package delete

import (
	"nfxid/modules/directory/domain/users"

	"gorm.io/gorm"
)

// Handler 处理删除操作，实现 users.Delete 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Delete Handler
func NewHandler(db *gorm.DB) users.Delete {
	return &Handler{db: db}
}
