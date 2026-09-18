package email

import (
	"nfxidentity/modules/auth/infrastructure/query/email/list"
	emailQuery "nfxidentity/modules/auth/query/email"

	"gorm.io/gorm"
)

func NewQuery(db *gorm.DB) *emailQuery.Query {
	return &emailQuery.Query{List: list.NewHandler(db)}
}
