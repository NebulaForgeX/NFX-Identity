package get

import (
	"nfxid/modules/access/domain/action_requirements"

	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) action_requirements.Get {
	return &Handler{db: db}
}
