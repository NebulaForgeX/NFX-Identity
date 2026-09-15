package phone

import (
	"nfxidentity/modules/auth/infrastructure/query/phone/list"
	phoneQuery "nfxidentity/modules/auth/query/phone"

	"gorm.io/gorm"
)

func NewQuery(db *gorm.DB) *phoneQuery.Query {
	return &phoneQuery.Query{List: list.NewHandler(db)}
}
