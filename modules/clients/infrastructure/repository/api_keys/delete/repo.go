package delete

import (
	"nfxid/modules/clients/domain/api_keys"

	"gorm.io/gorm"
)

// Handler 处理删除操作，实现 api_keys.Delete 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Delete Handler
func NewHandler(db *gorm.DB) api_keys.Delete {
	return &Handler{db: db}
}
