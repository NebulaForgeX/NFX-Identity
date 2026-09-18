package audios

import (
	"nfxidentity/modules/asset/infrastructure/query/audios/count"
	"nfxidentity/modules/asset/infrastructure/query/audios/list"
	"nfxidentity/modules/asset/infrastructure/query/audios/single"
	auds "nfxidentity/modules/asset/query/audios"

	"gorm.io/gorm"
)

func NewQuery(db *gorm.DB) *auds.Query {
	return &auds.Query{
		Single: single.NewHandler(db),
		List:   list.NewHandler(db),
		Count:  count.NewHandler(db),
	}
}
