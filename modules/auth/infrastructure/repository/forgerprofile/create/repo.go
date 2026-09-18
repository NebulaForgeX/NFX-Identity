package create

import (
	"nfxidentity/modules/auth/domain/forgerprofile"

	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) forgerprofile.Create { return &Handler{db: db} }
