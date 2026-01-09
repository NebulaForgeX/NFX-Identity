package check

import (
	"nfxid/modules/system/domain/system_state"

	"gorm.io/gorm"
)

// Handler 处理检查操作，实现 system_state.Check 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Check Handler
func NewHandler(db *gorm.DB) system_state.Check {
	return &Handler{db: db}
}
