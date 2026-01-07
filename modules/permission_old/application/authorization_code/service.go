package authorization_code

import (
	authorizationCodeDomain "nfxid/modules/permission/domain/authorization_code"
)

type Service struct {
	authorizationCodeRepo *authorizationCodeDomain.Repo
}

func NewService(
	authorizationCodeRepo *authorizationCodeDomain.Repo,
) *Service {
	return &Service{
		authorizationCodeRepo: authorizationCodeRepo,
	}
}
