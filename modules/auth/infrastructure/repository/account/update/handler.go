package update

import (
	accountDomain "nfxidentity/modules/auth/domain/account"

	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) accountDomain.Update { return &Handler{db: db} }
