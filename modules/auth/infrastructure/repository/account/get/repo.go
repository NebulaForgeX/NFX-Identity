package get

import (
	"nfxidentity/modules/auth/domain/account"
	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) account.Get { return &Handler{db: db} }
