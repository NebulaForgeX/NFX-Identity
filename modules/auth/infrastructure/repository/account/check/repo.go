package check

import (
	"nfxidentity/modules/auth/domain/account"
	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) account.Check { return &Handler{db: db} }
