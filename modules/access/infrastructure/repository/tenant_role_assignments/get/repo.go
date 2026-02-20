package get

import (
	"nfxid/modules/access/domain/tenant_role_assignments"
	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) tenant_role_assignments.Get {
	return &Handler{db: db}
}
