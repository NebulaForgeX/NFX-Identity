package image

import (
	imageDomain "nfxid/modules/image/domain/image"
	"nfxid/modules/image/infrastructure/query/image/list"
	"nfxid/modules/image/infrastructure/query/image/single"

	"gorm.io/gorm"
)

// NewHandler 创建新的 Image Query Handler
func NewHandler(db *gorm.DB) *imageDomain.Query {
	return &imageDomain.Query{
		Single: single.NewHandler(db),
		List:   list.NewHandler(db),
	}
}
