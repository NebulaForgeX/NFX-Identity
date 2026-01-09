package create

import (
	"nfxid/modules/clients/domain/client_scopes"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 client_scopes.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) client_scopes.Create {
	return &Handler{db: db}
}
