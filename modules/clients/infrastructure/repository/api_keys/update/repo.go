package update

import (
	"nfxid/modules/clients/domain/api_keys"

	"gorm.io/gorm"
)

// Handler 处理更新操作，实现 api_keys.Update 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Update Handler
func NewHandler(db *gorm.DB) api_keys.Update {
	return &Handler{db: db}
}
