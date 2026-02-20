package check

import (
	dom "nfxid/modules/access/domain/tenant_role_assignments"
	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) dom.Check {
	return &Handler{db: db}
}
