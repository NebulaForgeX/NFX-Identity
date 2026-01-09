package create

import (
	"nfxid/modules/tenants/domain/domain_verifications"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 domain_verifications.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) domain_verifications.Create {
	return &Handler{db: db}
}
