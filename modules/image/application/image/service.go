package image

import (
	imageQueries "nebulaid/modules/image/application/image/queries"
	imageDomain "nebulaid/modules/image/domain/image"
	"nebulaid/pkgs/eventbus"
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
