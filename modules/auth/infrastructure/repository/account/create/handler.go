package create

import (
	accountDomain "nfxidentity/modules/auth/domain/account"

	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) accountDomain.Create { return &Handler{db: db} }
