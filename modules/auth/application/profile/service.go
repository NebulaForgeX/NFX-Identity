package profile

import (
	profileQueries "nebulaid/modules/auth/application/profile/queries"
	profileDomain "nebulaid/modules/auth/domain/profile"
	"nebulaid/modules/auth/infrastructure/grpcclient"
)

type Service struct {
	profileRepo     profileDomain.Repo
	profileQuery    profileQueries.ProfileQuery
	imageGRPCClient *grpcclient.ImageGRPCClient
}

func NewService(
	profileRepo profileDomain.Repo,
	profileQuery profileQueries.ProfileQuery,
	imageGRPCClient *grpcclient.ImageGRPCClient,
) *Service {
	return &Service{
		profileRepo:     profileRepo,
		profileQuery:    profileQuery,
		imageGRPCClient: imageGRPCClient,
	}
}
