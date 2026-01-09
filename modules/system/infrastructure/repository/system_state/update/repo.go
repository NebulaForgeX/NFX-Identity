package update

import (
	"nfxid/modules/system/domain/system_state"

	"gorm.io/gorm"
)

// Handler 处理更新操作，实现 system_state.Update 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Update Handler
func NewHandler(db *gorm.DB) system_state.Update {
	return &Handler{db: db}
}
