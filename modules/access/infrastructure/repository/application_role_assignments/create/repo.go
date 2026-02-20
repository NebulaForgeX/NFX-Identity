package create

import (
	"nfxid/modules/access/domain/application_role_assignments"
	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) application_role_assignments.Create {
	return &Handler{db: db}
}
