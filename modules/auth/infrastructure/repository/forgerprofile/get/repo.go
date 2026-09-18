package get

import (
	"nfxidentity/modules/auth/domain/forgerprofile"

	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) forgerprofile.Get { return &Handler{db: db} }
