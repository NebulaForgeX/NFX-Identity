package create

import (
	"nfxid/modules/access/domain/super_admins"
	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) super_admins.Create {
	return &Handler{db: db}
}
