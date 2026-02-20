package create

import (
	dom "nfxid/modules/access/domain/tenant_role_assignments"
	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) dom.Create {
	return &Handler{db: db}
}
