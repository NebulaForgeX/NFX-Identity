package delete

import (
	"nfxid/modules/access/domain/action_requirements"

	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) action_requirements.Delete {
	return &Handler{db: db}
}
