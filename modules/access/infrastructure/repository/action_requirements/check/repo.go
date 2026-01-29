package check

import (
	"nfxid/modules/access/domain/action_requirements"

	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) action_requirements.Check {
	return &Handler{db: db}
}
