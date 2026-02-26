package image_types

import (
	imageTypeDomain "nfxidentity/modules/image/domain/image_types"
)

type Service struct {
	imageTypeRepo *imageTypeDomain.Repo
}

func NewService(
	imageTypeRepo *imageTypeDomain.Repo,
) *Service {
	return &Service{
		imageTypeRepo: imageTypeRepo,
	}
}
