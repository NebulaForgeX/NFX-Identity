package get

import (
	accountDomain "nfxidentity/modules/auth/domain/account"

	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) accountDomain.Get { return &Handler{db: db} }
