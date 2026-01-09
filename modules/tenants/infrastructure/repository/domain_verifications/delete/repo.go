package delete

import (
	"nfxid/modules/tenants/domain/domain_verifications"

	"gorm.io/gorm"
)

// Handler 处理删除操作，实现 domain_verifications.Delete 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Delete Handler
func NewHandler(db *gorm.DB) domain_verifications.Delete {
	return &Handler{db: db}
}
