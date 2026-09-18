package update

import (
	"nfxidentity/modules/auth/domain/forgerprofile"

	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) forgerprofile.Update { return &Handler{db: db} }
