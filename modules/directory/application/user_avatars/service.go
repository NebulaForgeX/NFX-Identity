package user_avatars

import (
	"nfxid/connections/image"
	userAvatarDomain "nfxid/modules/directory/domain/user_avatars"
	"nfxid/pkgs/kafkax/eventbus"
)

type Service struct {
	userAvatarRepo *userAvatarDomain.Repo
	imageClient    *image.ImageClient // create/update 前通过 gRPC 校验 image 存在、移动 tmp→avatar
	busPublisher   *eventbus.BusPublisher
}

func NewService(
	userAvatarRepo *userAvatarDomain.Repo,
	imageClient *image.ImageClient,
	busPublisher *eventbus.BusPublisher,
) *Service {
	return &Service{
		userAvatarRepo: userAvatarRepo,
		imageClient:    imageClient,
		busPublisher:   busPublisher,
	}
}
