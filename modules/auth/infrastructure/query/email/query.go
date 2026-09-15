package email

import (
	emailQuery "nfxidentity/modules/auth/query/email"
	"nfxidentity/modules/auth/infrastructure/query/email/list"

	"gorm.io/gorm"
)

func NewQuery(db *gorm.DB) *emailQuery.Query {
	return &emailQuery.Query{List: list.NewHandler(db)}
}
