package images

import (
	imageDomain "nfxid/modules/image/domain/images"
)

type Service struct {
	imageRepo *imageDomain.Repo
}

func NewService(
	imageRepo *imageDomain.Repo,
) *Service {
	return &Service{
		imageRepo: imageRepo,
	}
}
