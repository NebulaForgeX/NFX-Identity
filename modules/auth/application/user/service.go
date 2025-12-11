package user

import (
	"nfxid/modules/auth/application/user/queries"
	profileDomain "nfxid/modules/auth/domain/profile"
	userDomain "nfxid/modules/auth/domain/user"
	"nfxid/pkgs/eventbus"
	"nfxid/pkgs/tokenx"
)

type Service struct {
	userRepo     userDomain.Repo
	profileRepo  profileDomain.Repo
	userQuery    queries.UserQuery
	busPublisher *eventbus.BusPublisher
	tokenx       *tokenx.Tokenx
}

func NewService(
	userRepo userDomain.Repo,
	profileRepo profileDomain.Repo,
	userQuery queries.UserQuery,
	busPublisher *eventbus.BusPublisher,
	tokenx *tokenx.Tokenx,
) *Service {
	return &Service{
		userRepo:     userRepo,
		profileRepo:  profileRepo,
		userQuery:    userQuery,
		busPublisher: busPublisher,
		tokenx:       tokenx,
	}
}
