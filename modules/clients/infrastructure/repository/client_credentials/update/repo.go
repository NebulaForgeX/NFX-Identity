package update

import (
	"nfxid/modules/clients/domain/client_credentials"

	"gorm.io/gorm"
)

// Handler 处理更新操作，实现 client_credentials.Update 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Update Handler
func NewHandler(db *gorm.DB) client_credentials.Update {
	return &Handler{db: db}
}
