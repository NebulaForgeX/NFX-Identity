package create

import (
	"nfxid/modules/access/domain/actions"

	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) actions.Create {
	return &Handler{db: db}
}
