package delete

import (
	"nfxid/modules/access/domain/tenant_role_assignments"
	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) tenant_role_assignments.Delete {
	return &Handler{db: db}
}
