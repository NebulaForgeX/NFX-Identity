package update

import (
	"nfxid/modules/tenants/domain/domain_verifications"

	"gorm.io/gorm"
)

// Handler 处理更新操作，实现 domain_verifications.Update 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Update Handler
func NewHandler(db *gorm.DB) domain_verifications.Update {
	return &Handler{db: db}
}
