package create

import (
	"nfxid/modules/access/domain/application_roles"
	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) application_roles.Create {
	return &Handler{db: db}
}
