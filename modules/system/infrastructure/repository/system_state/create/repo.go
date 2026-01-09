package create

import (
	"nfxid/modules/system/domain/system_state"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 system_state.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) system_state.Create {
	return &Handler{db: db}
}
