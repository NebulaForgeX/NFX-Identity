package create

import (
	"nfxid/modules/access/domain/action_requirements"

	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) action_requirements.Create {
	return &Handler{db: db}
}
