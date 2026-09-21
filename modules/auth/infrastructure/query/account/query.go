package account

import (
	"nfxidentity/modules/auth/infrastructure/query/account/single"
	accountQuery "nfxidentity/modules/auth/query/account"

	"gorm.io/gorm"
)

func NewQuery(db *gorm.DB) *accountQuery.Query {
	return &accountQuery.Query{
		Single: single.NewHandler(db),
	}
}
