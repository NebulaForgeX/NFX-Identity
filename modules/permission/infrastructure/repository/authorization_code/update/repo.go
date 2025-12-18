package update

import (
	authorizationCodeDomain "nfxid/modules/permission/domain/authorization_code"

	"gorm.io/gorm"
)

// Handler 处理更新操作，实现 authorizationCodeDomain.Update 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Update Handler
func NewHandler(db *gorm.DB) authorizationCodeDomain.Update {
	return &Handler{db: db}
}
