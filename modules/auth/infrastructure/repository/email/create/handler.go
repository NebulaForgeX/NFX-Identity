package create

import (
	emailDomain "nfxidentity/modules/auth/domain/email"

	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) emailDomain.Create { return &Handler{db: db} }
