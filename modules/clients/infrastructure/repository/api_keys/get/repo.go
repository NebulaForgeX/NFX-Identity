package get

import (
	"nfxid/modules/clients/domain/api_keys"

	"gorm.io/gorm"
)

// Handler 处理获取数据操作，实现 api_keys.Get 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Get Handler
func NewHandler(db *gorm.DB) api_keys.Get {
	return &Handler{db: db}
}
