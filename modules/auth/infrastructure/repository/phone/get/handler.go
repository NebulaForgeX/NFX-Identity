package get

import (
	phoneDomain "nfxidentity/modules/auth/domain/phone"

	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) phoneDomain.Get { return &Handler{db: db} }
