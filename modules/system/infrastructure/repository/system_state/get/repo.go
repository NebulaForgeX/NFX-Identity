package get

import (
	"nfxid/modules/system/domain/system_state"

	"gorm.io/gorm"
)

// Handler 处理获取数据操作，实现 system_state.Get 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Get Handler
func NewHandler(db *gorm.DB) system_state.Get {
	return &Handler{db: db}
}
