package user_images

import (
	"nfxid/connections/image"
	userImageDomain "nfxid/modules/directory/domain/user_images"
)

type Service struct {
	userImageRepo *userImageDomain.Repo
	imageClient   *image.ImageClient // 可选：create/update 前通过 gRPC 问 Image 服务该 image 是否存在
}

func NewService(
	userImageRepo *userImageDomain.Repo,
	imageClient *image.ImageClient,
) *Service {
	return &Service{
		userImageRepo: userImageRepo,
		imageClient:   imageClient,
	}
}
