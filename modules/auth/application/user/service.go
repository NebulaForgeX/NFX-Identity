package user

import (
	userDomain "nfxid/modules/auth/domain/user"
	"nfxid/pkgs/eventbus"
	"nfxid/pkgs/tokenx"
)

type Service struct {
	userRepo     *userDomain.Repo
	userQuery    *userDomain.Query
	busPublisher *eventbus.BusPublisher
	tokenx       *tokenx.Tokenx
}

func NewService(
	userRepo *userDomain.Repo,
	userQuery *userDomain.Query,
	busPublisher *eventbus.BusPublisher,
	tokenx *tokenx.Tokenx,
) *Service {
	return &Service{
		userRepo:     userRepo,
		userQuery:    userQuery,
		busPublisher: busPublisher,
		tokenx:       tokenx,
	}
}
