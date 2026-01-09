package create

import (
	"nfxid/modules/clients/domain/api_keys"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 api_keys.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) api_keys.Create {
	return &Handler{db: db}
}
