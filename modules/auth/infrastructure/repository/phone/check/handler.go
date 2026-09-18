package check

import (
	phoneDomain "nfxidentity/modules/auth/domain/phone"

	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) phoneDomain.Check { return &Handler{db: db} }
