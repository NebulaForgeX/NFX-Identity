package update

import (
	"nfxid/modules/clients/domain/client_scopes"

	"gorm.io/gorm"
)

// Handler 处理更新操作，实现 client_scopes.Update 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Update Handler
func NewHandler(db *gorm.DB) client_scopes.Update {
	return &Handler{db: db}
}
