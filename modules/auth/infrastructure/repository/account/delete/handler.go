package delete

import (
	accountDomain "nfxidentity/modules/auth/domain/account"

	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) accountDomain.Delete { return &Handler{db: db} }
