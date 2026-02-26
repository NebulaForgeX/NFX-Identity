package image_tags

import (
	imageTagDomain "nfxidentity/modules/image/domain/image_tags"
)

type Service struct {
	imageTagRepo *imageTagDomain.Repo
}

func NewService(
	imageTagRepo *imageTagDomain.Repo,
) *Service {
	return &Service{
		imageTagRepo: imageTagRepo,
	}
}
