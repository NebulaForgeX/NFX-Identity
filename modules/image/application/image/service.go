package image

import (
	imageQueries "nfxid/modules/image/application/image/queries"
	imageDomain "nfxid/modules/image/domain/image"
	"nfxid/pkgs/eventbus"
)

type Service struct {
	imageRepo    imageDomain.Repo
	imageQuery   imageQueries.ImageQuery
	busPublisher *eventbus.BusPublisher
}

func NewService(
	imageRepo imageDomain.Repo,
	imageQuery imageQueries.ImageQuery,
	busPublisher *eventbus.BusPublisher,
) *Service {
	return &Service{
		imageRepo:    imageRepo,
		imageQuery:   imageQuery,
		busPublisher: busPublisher,
	}
}
