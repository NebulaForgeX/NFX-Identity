package get

import (
	"nfxid/modules/access/domain/super_admins"
	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) super_admins.Get {
	return &Handler{db: db}
}
