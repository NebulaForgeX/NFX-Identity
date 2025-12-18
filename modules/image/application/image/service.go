package image

import (
	imageDomain "nfxid/modules/image/domain/image"
	"nfxid/pkgs/eventbus"
)

type Service struct {
	imageRepo    *imageDomain.Repo
	imageQuery   *imageDomain.Query
	busPublisher *eventbus.BusPublisher
}

func NewService(
	imageRepo *imageDomain.Repo,
	imageQuery *imageDomain.Query,
	busPublisher *eventbus.BusPublisher,
) *Service {
	return &Service{
		imageRepo:    imageRepo,
		imageQuery:   imageQuery,
		busPublisher: busPublisher,
	}
}
