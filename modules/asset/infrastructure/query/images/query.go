package images

import (
	"nfxidentity/modules/asset/infrastructure/query/images/single"
	imagesQuery "nfxidentity/modules/asset/query/images"

	"gorm.io/gorm"
)

func NewQuery(db *gorm.DB) *imagesQuery.Query {
	return &imagesQuery.Query{Single: single.NewHandler(db)}
}
