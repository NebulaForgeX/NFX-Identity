package user_images

import (
	userImageDomain "nfxid/modules/directory/domain/user_images"
)

type Service struct {
	userImageRepo *userImageDomain.Repo
}

func NewService(
	userImageRepo *userImageDomain.Repo,
) *Service {
	return &Service{
		userImageRepo: userImageRepo,
	}
}
