package get

import (
	"nfxid/modules/clients/domain/client_scopes"

	"gorm.io/gorm"
)

// Handler 处理获取数据操作，实现 client_scopes.Get 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Get Handler
func NewHandler(db *gorm.DB) client_scopes.Get {
	return &Handler{db: db}
}
