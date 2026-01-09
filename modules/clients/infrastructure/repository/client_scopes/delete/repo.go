package delete

import (
	"nfxid/modules/clients/domain/client_scopes"

	"gorm.io/gorm"
)

// Handler 处理删除操作，实现 client_scopes.Delete 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Delete Handler
func NewHandler(db *gorm.DB) client_scopes.Delete {
	return &Handler{db: db}
}
