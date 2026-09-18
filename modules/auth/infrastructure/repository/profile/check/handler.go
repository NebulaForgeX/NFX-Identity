package check

import (
	profileDomain "nfxidentity/modules/auth/domain/profile"

	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) profileDomain.Check {
	return &Handler{db: db}
}
