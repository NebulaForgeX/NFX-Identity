package delete

import (
	phoneDomain "nfxidentity/modules/auth/domain/phone"

	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) phoneDomain.Delete { return &Handler{db: db} }
