package delete

import (
	"nfxid/modules/access/domain/actions"

	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) actions.Delete {
	return &Handler{db: db}
}
