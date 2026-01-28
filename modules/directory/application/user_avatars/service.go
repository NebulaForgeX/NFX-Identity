package user_avatars

import (
	userAvatarDomain "nfxid/modules/directory/domain/user_avatars"
)

type Service struct {
	userAvatarRepo *userAvatarDomain.Repo
}

func NewService(
	userAvatarRepo *userAvatarDomain.Repo,
) *Service {
	return &Service{
		userAvatarRepo: userAvatarRepo,
	}
}
