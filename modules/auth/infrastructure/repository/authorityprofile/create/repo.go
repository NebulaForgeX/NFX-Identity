package create

import (
	"nfxidentity/modules/auth/domain/authorityprofile"

	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) authorityprofile.Create { return &Handler{db: db} }
