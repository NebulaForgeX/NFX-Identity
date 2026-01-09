package delete

import (
	"nfxid/modules/system/domain/system_state"

	"gorm.io/gorm"
)

// Handler 处理删除操作，实现 system_state.Delete 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Delete Handler
func NewHandler(db *gorm.DB) system_state.Delete {
	return &Handler{db: db}
}
