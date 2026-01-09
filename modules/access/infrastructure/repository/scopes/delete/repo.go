package delete

import (
	"nfxid/modules/access/domain/scopes"

	"gorm.io/gorm"
)

// Handler 处理删除操作，实现 scopes.Delete 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Delete Handler
func NewHandler(db *gorm.DB) scopes.Delete {
	return &Handler{db: db}
}
