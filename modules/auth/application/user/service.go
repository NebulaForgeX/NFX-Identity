package user

import (
	"nebulaid/modules/auth/application/user/queries"
	profileDomain "nebulaid/modules/auth/domain/profile"
	userDomain "nebulaid/modules/auth/domain/user"
	"nebulaid/pkgs/eventbus"
	"nebulaid/pkgs/tokenx"
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
