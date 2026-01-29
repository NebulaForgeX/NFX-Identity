package get

import (
	"nfxid/modules/access/domain/actions"

	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) actions.Get {
	return &Handler{db: db}
}
