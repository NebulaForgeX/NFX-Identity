package update

import (
	"nfxid/modules/access/domain/actions"

	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) actions.Update {
	return &Handler{db: db}
}
