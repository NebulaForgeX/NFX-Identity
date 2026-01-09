package check

import (
	"nfxid/modules/tenants/domain/domain_verifications"

	"gorm.io/gorm"
)

// Handler 处理检查操作，实现 domain_verifications.Check 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Check Handler
func NewHandler(db *gorm.DB) domain_verifications.Check {
	return &Handler{db: db}
}
