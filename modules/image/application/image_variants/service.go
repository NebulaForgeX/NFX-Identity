package image_variants

import (
	imageVariantDomain "nfxid/modules/image/domain/image_variants"
)

type Service struct {
	imageVariantRepo *imageVariantDomain.Repo
}

func NewService(
	imageVariantRepo *imageVariantDomain.Repo,
) *Service {
	return &Service{
		imageVariantRepo: imageVariantRepo,
	}
}
